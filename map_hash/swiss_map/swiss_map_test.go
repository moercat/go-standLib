package swiss_map

import "testing"

func BenchmarkName(b *testing.B) {
	Init()
	for i := 0; i < b.N; i++ {
		GetMap()
	}
}
