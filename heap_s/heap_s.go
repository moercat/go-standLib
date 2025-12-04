package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"bt.baishancloud.com/mtrpc/pb/auth"
	"bt.baishancloud.com/mtrpc/security/author"
)

const (
	ttPredBill            = 4.2 * 1.123
	configPath            = "/data/config/config_pstatp_p.json"
	apiURL                = "https://dragon-api.bs58i.baishancdnx.com/traffic/topn_users_detail_traffic/%d/%d/1/flow"
	rollbackAPIURL        = "https://water-test.ort.sealaly.com/v-api/strategy-distribute/update"
	wechatBotURL          = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + AbnormalAlarmRobotKey
	heapPayload           = `{"id":705,"strategy_config_id":4167,"remark":"关闭下载控线，如需回退，回退到3925"}`
	rollbackPayload       = `{"id":705,"strategy_config_id":3925,"remark":"如需控线 4167"}`
	AbnormalAlarmRobotKey = "d9775657-4e75-4188-a094-d326d9ee7a48"
	ZttUserID             = "29013"
)

var (
	authKeeper author.AuthKeeper // 非指针类型
	heapStatus string
	heapMutex  sync.RWMutex
)

type Config struct {
	IsHeap string `json:"是否堆高"`
}
type TrafficData struct {
	Meta struct {
		RequestId string `json:"request_id"`
		Code      int    `json:"code"`
		Error     string `json:"error"`
	} `json:"meta"`
	Err  string                   `json:"err"`
	Data []map[string][][]float64 `json:"data"`
}
type WechatMessage struct {
	MsgType  string `json:"msgtype"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}

type timeBw struct {
	Time      string
	Bandwidth float64
}
type ReqInfo struct {
	Type    string   `json:"type"`
	UserIds []string `json:"user_ids"`
}

func main() {
	os.Setenv("MTRPC_PROXY_PRO_JH", "digital-rpcproxy.bs58i.baishancdnx.com:80")
	os.Setenv("DATACENTER", "local")
	proxy := os.Getenv("MTRPC_PROXY_PRO_JH")
	datacenter := os.Getenv("DATACENTER")
	fmt.Printf("当前数据中心:proxy %s datacenter%s\n", proxy, datacenter)
	appid, secret := "bsms3466186", "cwpzpv74qkovn96487c7"
	// 初始化认证管理器（非指针）
	authKeeper = author.MustSetupAuthKeeper(appid, secret,
		author.SetupAuthWithArgs(&auth.GenerateTokenV2Request{
			Scene: "in",
		}))
	token, err := authKeeper.GetToken()
	if err != nil {
		fmt.Printf("获取token失败: %v", err)
		return
	}
	initHeapStatus()

	now := time.Now()
	dfData, err := fetchTrafficData(now.Unix()-1200, now.Unix(), token)
	if err != nil {
		fmt.Printf("获取流量数据失败: %v\n", err)
		return
	}

	if len(dfData) < 2 {
		fmt.Println("错误：流量数据不足")
		return
	}

	recentBandwidth := dfData[len(dfData)-2].Bandwidth
	heap := ifHeap(recentBandwidth, ttPredBill)
	fmt.Printf("堆高状态: %v, 带宽: %.4fT, 阈值: %.4fT\n", heap, recentBandwidth, ttPredBill*1.06)

	handleHeapStatus(heap, recentBandwidth)
	saveHeapStatus()
}

func initHeapStatus() {
	config, err := readConfig()
	if err != nil {
		fmt.Printf("读取配置失败: %v, 使用默认值'否'\n", err)
		heapStatus = "否"
		return
	}
	heapStatus = config.IsHeap
	fmt.Printf("初始化堆高状态: %s\n", heapStatus)
}

func saveHeapStatus() {
	data, err := json.MarshalIndent(Config{IsHeap: heapStatus}, "", "    ")
	if err != nil {
		fmt.Printf("序列化配置失败: %v\n", err)
		return
	}
	if err := ioutil.WriteFile(configPath, data, 0644); err != nil {
		fmt.Printf("保存配置文件失败: %v\n", err)
	} else {
		fmt.Printf("堆高状态已保存: %s\n", heapStatus)
	}
}

func getHeapStatus() string {
	heapMutex.RLock()
	defer heapMutex.RUnlock()
	return heapStatus
}

func setHeapStatus(status string) {
	heapMutex.Lock()
	defer heapMutex.Unlock()
	heapStatus = status
}

func fetchTrafficData(start, end int64, token string) ([]timeBw, error) {
	var result []timeBw

	reqInfo := ReqInfo{
		Type: "edge",
		UserIds: []string{
			ZttUserID,
		}}

	reqInfoByte, err := json.Marshal(reqInfo)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(apiURL, start, end)
	fmt.Println(url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqInfoByte))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	fmt.Println(token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var trafficData TrafficData
	if err := json.Unmarshal(body, &trafficData); err != nil {
		return nil, err
	}
	if trafficData.Meta.Code != 0 {
		return nil, fmt.Errorf("API错误: %s", trafficData.Meta.Error)
	}
	if trafficData.Err != "" {
		return nil, fmt.Errorf("API错误: %s", trafficData.Err)
	}

	fmt.Println(trafficData)

	if len(trafficData.Data) > 0 {
		if userData, ok := trafficData.Data[0][ZttUserID]; ok {
			fmt.Println(userData)
			for _, dataPoint := range userData {
				if len(dataPoint) < 2 {
					continue
				}
				t := time.Unix(int64(dataPoint[0]), 0).Format("2006-01-02 15:04:05")
				bandwidth := dataPoint[1] * 8 / 300 / 1e12
				result = append(result, struct {
					Time      string
					Bandwidth float64
				}{Time: t, Bandwidth: bandwidth})
			}
		}
	}
	return result, nil
}

func ifHeap(bw, pred float64) bool { return bw >= pred*1.06 }

func handleHeapStatus(heap bool, bw float64) {
	currentStatus := getHeapStatus()
	if heap {
		if currentStatus == "是" {
			fmt.Println("堆高状态未变化")
			return
		}
		setHeapStatus("是")
		fmt.Println("堆高状态已更新为'是'")

		if err := triggerRollback(heapPayload); err != nil {
			fmt.Printf("触发关闭下载控线失败: %v\n", err)
		} else {
			fmt.Println("已触发关闭下载控线")
		}
		sendWechatMessage("头条新账号pstatp-p要堆高\n当前带宽: **" + strconv.FormatFloat(bw, 'f', 2, 64) + "T**")
	} else {
		if currentStatus == "否" {
			fmt.Println("非堆高状态未变化")
			return
		}
		setHeapStatus("否")
		fmt.Println("堆高状态已更新为'否'")

		if err := triggerRollback(rollbackPayload); err != nil {
			fmt.Printf("触发回退到3925失败: %v\n", err)
		} else {
			fmt.Println("已触发回退到3925")
		}
		sendWechatMessage("头条新账号pstatp-p堆高恢复\n当前带宽: **" + strconv.FormatFloat(bw, 'f', 2, 64) + "T**")
	}
}

func readConfig() (Config, error) {
	var config Config
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return Config{IsHeap: "否"}, nil
	}
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config, err
	}
	if err := json.Unmarshal(file, &config); err != nil {
		return config, err
	}
	return config, nil
}

func triggerRollback(payload string) error {
	req, err := http.NewRequest("POST", rollbackAPIURL, bytes.NewBufferString(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "dfebbda7af1818533fe8a1d00f8e98724fdaa809")
	req.Header.Set("Authorization-Type", "apikey")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("API响应: %s\n", string(body))
	return nil
}

func sendWechatMessage(content string) {
	msg := WechatMessage{MsgType: "markdown"}
	msg.Markdown.Content = content
	jsonData, _ := json.Marshal(msg)

	req, _ := http.NewRequest("POST", wechatBotURL, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("发送企业微信通知失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("企业微信API错误: %s, 响应: %s\n", resp.Status, string(body))
	} else {
		fmt.Println("已发送企业微信通知")
	}
}
