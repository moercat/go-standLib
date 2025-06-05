package if_var

import "testing"

func BenchmarkNumber1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Number1()
	}
}

func BenchmarkNumber2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Number2()
	}
}
