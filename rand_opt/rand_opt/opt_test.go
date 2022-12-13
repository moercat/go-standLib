package opt

import (
	"fmt"
	"testing"
)

func BenchmarkName_old(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Old()
	}
}

func BenchmarkName_Rng(b *testing.B) {
	var r RNG
	for i := 0; i < b.N; i++ {
		Rng(r)
	}
}

func Test_Rand(t *testing.T) {
	var (
		r RNG
		m = make(map[uint64]int, 0)
	)

	for i := 0; i < 10000000; i++ {
		n := r.Uint64n(10000)
		m[n]++
	}

	fmt.Println(m)

}
