package opt

import (
	"math/rand"
)

func MathRand() {
	go func() {
		for i := 0; i < 100; i++ {
			rand.Intn(10)
		}
	}()
}

func Fastrand() {
	go func() {
		for i := 0; i < 100; i++ {
			Fastrandn(10)
		}
	}()
}

func MutexRand(o RandWithMutex) {
	go func(o RandWithMutex) {
		for i := 0; i < 100; i++ {
			o.Lock.Lock()
			o.R.Intn(10)
			o.Lock.Unlock()
		}
	}(o)
}

func RWRand(o RandWithRW) {
	go func(o RandWithRW) {
		for i := 0; i < 100; i++ {
			o.Lock.Lock()
			o.R.Intn(10)
			o.Lock.Unlock()
		}
	}(o)
}

func Rng(r RNG) {
	go func(r RNG) {
		for i := 0; i < 100; i++ {
			r.Uint64n(10)
		}
	}(r)
}
