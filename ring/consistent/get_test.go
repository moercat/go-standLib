package consistent

import (
	"testing"
)

var (
	key = "111111111111"
)

func BenchmarkInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Init()
	}
}

func BenchmarkInitStr(b *testing.B) {

	for i := 0; i < b.N; i++ {
		InitStr()
	}
}

func BenchmarkGet(b *testing.B) {
	Init()
	for i := 0; i < b.N; i++ {
		Get(key)
	}
}

func BenchmarkGetStr(b *testing.B) {
	InitStr()
	for i := 0; i < b.N; i++ {
		GetStr(key)
	}
}
