package map1

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func BenchmarkName1(b *testing.B) {
	n := time.Now()
	for i := 0; i < 100000000; i++ {
		EEE()
	}
	fmt.Println(time.Since(n))
}

func BenchmarkName2(b *testing.B) {
	n := time.Now()
	for i := 0; i < 100000000; i++ {
		FFF()
	}
	fmt.Println(time.Since(n))
}

var bizMap sync.Map

func TestName3(t *testing.T) {

	var map111 map[string]map[string]string

	a := map111["s"]

	fmt.Println(a["s"])
}

// parseBiz 解析业务相关到BizTypeMap
func parseBiz(bizTypes, bizParams string) (bizTypeMapResp map[string][]string) {
	if value, ok := bizMap.Load(bizTypes + "::" + bizParams); ok {
		items := value.(map[string][]string)
		bizTypeMapResp = make(map[string][]string, len(items))
		for k, v := range items {
			bizTypeMapResp[k] = v
		}
		return
	}
	bizTypeMap := make(map[string][]string)
	bizTypeMap["aaa"] = []string{"a", "a", "a"}
	bizTypeMap["bbb"] = []string{"b", "b", "b"}
	bizTypeMap["ccc"] = []string{"c", "c", "c"}
	bizTypeMap[bizParams] = []string{bizTypes, bizParams, "::"}

	fmt.Println("store", bizTypeMap)
	bizMap.Store(bizTypes+"::"+bizParams, bizTypeMap)
	bizTypeMapResp = make(map[string][]string, len(bizTypeMap))
	for k, v := range bizTypeMap {
		bizTypeMapResp[k] = v
	}
	return
}
