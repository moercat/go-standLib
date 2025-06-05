package fast_search

import (
	"testing"
)

func init() {
	Hash()
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get()
	}
}

func BenchmarkGetWithIdx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetWithIdx()
	}
}

func TestGet(t *testing.T) {
	//Hash()
	for i := 0; i < 1; i++ {
		Get()
	}
	//time.Sleep(time.Second)
	for i := 0; i < 1; i++ {
		GetWithIdx()
	}
	//time.Sleep(time.Minute)
}

var mm = map[string]string{
	"aa": "aa",
	"bb": "bb",
	"cc": "cc",
	"dd": "dd",
}

type strr struct {
	Key   string
	Value string
}

var ll = []strr{
	{
		Key:   "aa",
		Value: "aa",
	},
	{
		Key:   "bb",
		Value: "bb",
	},
	{
		Key:   "cc",
		Value: "cc",
	},
	{
		Key:   "dd",
		Value: "dd",
	},
}

var (
	str1 string
	ok1  bool
)

func BenchmarkMM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str, ok := mm["cc"]
		str1 = str
		ok1 = ok
		//fmt.Println(str, ok)
	}
}

func BenchmarkLL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, s := range ll {
			if s.Key == "cc" {
				str1 = s.Key
				ok1 = true
			}
		}
	}
}
