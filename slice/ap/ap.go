package ap

import (
	"fmt"
	"sort"
)

//func main() {
//
//	temp := make(map[string]int)
//	fmt.Println("len(temp):", len(temp), "temp:", temp)
//
//	temp1 := make(map[string]int, 3)
//	fmt.Println("len(temp1):", len(temp1), "temp1", temp1)
//
//}

func add(sNodeName []string) {
	for i := 0; i < 4; i++ {
		sNodeName = append(sNodeName, "11111")
	}
}

type SdkAvailableHunter struct {
	Isp           string `json:"isp,omitempty"`             //运营商
	Region        string `json:"region,omitempty"`          // 大区
	Province      string `json:"province,omitempty"`        // 省份
	HunterGroup   string `json:"hunter_group,omitempty"`    // hunter 组
	HunterSvrName string `json:"hunter_svr_name,omitempty"` // hunter 设备名
	HunterIp      string `json:"hunter_ip,omitempty"`       // hunter Ip
}

func ListCheck() {
	var availableHunters = []SdkAvailableHunter{
		{
			Isp:      "dx",
			Region:   "huadong",
			Province: "fujian",
		},
		{
			Isp:      "lt",
			Region:   "huadong",
			Province: "fujian",
		},
		{
			Isp:      "yd",
			Region:   "huadong",
			Province: "fujian",
		},
		{
			Isp:      "dx",
			Region:   "huadong",
			Province: "anhui",
		},
		{
			Isp:      "yd",
			Region:   "xinan",
			Province: "yunnan",
		},
		{
			Isp:      "yd",
			Region:   "xinan",
			Province: "sichuan",
		},
	}

	var b []SdkAvailableHunter
	for _, hunter := range availableHunters {
		b = append(b, hunter)
	}

	for _, group := range b {
		fmt.Println(group)
		sort.Slice(availableHunters, func(i, j int) bool {
			var (
				iIdx, jIdx = 10, 10
			)
			if availableHunters[i].Region == group.Region {
				iIdx = 4
			}
			if availableHunters[j].Region == group.Region {
				jIdx = 4
			}
			if availableHunters[i].Province == group.Province {
				iIdx = 3
			}
			if availableHunters[j].Province == group.Province {
				jIdx = 3
			}
			if availableHunters[i].Isp == group.Isp && availableHunters[i].Region == group.Region {
				iIdx = 2
			}
			if availableHunters[j].Isp == group.Isp && availableHunters[j].Region == group.Region {
				jIdx = 2
			}
			if availableHunters[i].Isp == group.Isp && availableHunters[i].Province == group.Province {
				iIdx = 1
			}
			if availableHunters[j].Isp == group.Isp && availableHunters[j].Province == group.Province {
				jIdx = 1
			}
			return iIdx < jIdx
		})
		fmt.Println(group, availableHunters)
	}

}
