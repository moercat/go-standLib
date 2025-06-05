package spilt

import (
	"fmt"
	"testing"
)

func TestStr(t *testing.T) {
	fmt.Println(S(), SN())
}

func BenchmarkSNName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		S()
	}
}

func BenchmarkSName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SN()
	}
}

func Test111(t *testing.T) {
	var toRemoteCount = int(0)
	var max = 1

	for i := 0; i < 5; i++ {
		fofff(&toRemoteCount, max)
	}

}

func fofff(toRemoteCount *int, max int) {
	for i := 0; i < 10; i++ {
		if *toRemoteCount < max {
			*toRemoteCount = *toRemoteCount + 1
			fmt.Println(i, *toRemoteCount)
		}
	}
}

func TestC(t *testing.T) {
	fmt.Println(C(), CR())
}

func BenchmarkCName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		C()
	}
}

func BenchmarkCRName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CR()
	}
}
