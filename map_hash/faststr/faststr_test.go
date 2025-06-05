package faststr

import "testing"

func BenchmarkPMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PMap()
	}
}

func BenchmarkPSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PSlice()
	}
}
