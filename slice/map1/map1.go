package map1

import (
	"errors"
	"fmt"
	"strconv"
	"sync/atomic"
)

var list []StrMsg
var set = make(map[string]string)

type StrMsg struct {
	Key   string
	Value string
}

var (
	moreMap = make(map[string]*int64)
	lessMap = make(map[string]*int64)
)

func init() {

	for i := 0; i < 5; i++ {
		j := i + 100
		list = append(list, StrMsg{
			Key:   strconv.Itoa(j),
			Value: strconv.Itoa(j),
		})
		set[strconv.Itoa(j)] = strconv.Itoa(j)
	}

	for i := 0; i < 10000; i++ {
		moreMap[strconv.Itoa(i)+"qqqqqqqqqqqqqqqq"] = new(int64)
		if i > 5000 {
			continue
		}
		lessMap[strconv.Itoa(i)+"qqqqqqqqqqqqqqqq"] = new(int64)
	}

	var m map[string]string

	m["2"] = "2"

}

func AAA(tag string) (string, error) {
	for _, v := range list {
		if v.Key == tag {
			return v.Value, nil
		}
	}
	return "", errors.New("aaa")
}

func BBB(tag string) (string, error) {
	value, has := set[tag]
	if !has {
		return "", errors.New("aaa")
	}
	return value, nil
}

func CCC() {

	var a = make([]int, 1, 10)

	DDD(a)

	fmt.Println(a[2])

}

func DDD(aaa []int) []int {
	return append(aaa, 1)
}

func EEE() {
	v, exist := moreMap["6"+"qqqqqqqqqqqqqqqq"]
	if !exist {
		return
	}

	atomic.AddInt64(v, 1)
}

func FFF() {
	v, exist := lessMap["6"+"qqqqqqqqqqqqqqqq"]
	if !exist {
		return
	}

	atomic.AddInt64(v, 1)
}
