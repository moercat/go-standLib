package two_map

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
)

func TestNumber(t *testing.T) {

	for i := 0; i < 100000; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				Number(rand.Intn(3000) % 1000)
			}
		}()

		go func(i int) {
			for j := 0; j < 1000; j++ {
				Number(i)
			}
		}(i)
	}

	for _, m := range fenceAllMap {
		fmt.Println(atomic.LoadInt64(m[1]))
	}
}
