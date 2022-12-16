package opt

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"
)

func BenchmarkName_mathrand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MathRand()
	}
}

func BenchmarkName_fast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fastrand()
	}
}

func BenchmarkName_mutex(b *testing.B) {

	var r = RandWithMutex{
		R:    rand.New(rand.NewSource(time.Now().UnixNano())),
		Lock: &sync.Mutex{},
	}
	for i := 0; i < b.N; i++ {

		MutexRand(r)
	}
}

func BenchmarkName_rw(b *testing.B) {

	var r = RandWithRW{
		R:    rand.New(rand.NewSource(time.Now().UnixNano())),
		Lock: &sync.RWMutex{},
	}
	for i := 0; i < b.N; i++ {

		RWRand(r)
	}
}

func Test_mathRand(t *testing.T) {
	var (
		l = make([]int, 1000)
	)

	for i := 0; i < 10000000; i++ {
		n := rand.Intn(1000)
		l[n]++
	}

	time.Sleep(time.Second)

	sort.Ints(l)

	fmt.Println(l[999] - l[0])
}

func Test_Fast(t *testing.T) {
	var (
		l = make([]int, 1000)
	)

	for i := 0; i < 10000000; i++ {
		n := Fastrandn(1000)
		l[n]++
	}

	time.Sleep(time.Second)

	sort.Ints(l)

	fmt.Println(l[999] - l[0])
}

func Test_Mutex(t *testing.T) {

	var l = make([]int, 1000)

	var r = RandWithMutex{
		R:    rand.New(rand.NewSource(time.Now().UnixNano())),
		Lock: &sync.Mutex{},
	}

	for i := 0; i < 10000000; i++ {

		go func(o RandWithMutex) {
			for i := 0; i < 100; i++ {
				o.Lock.Lock()
				n := o.R.Intn(1000)
				o.Lock.Unlock()
				l[n]++
			}
		}(r)
	}

	time.Sleep(time.Second)

	sort.Ints(l)

	fmt.Println(l[999] - l[0])

}

func Test_RW(t *testing.T) {

	var l = make([]int, 1000)

	var r = RandWithRW{
		R:    rand.New(rand.NewSource(time.Now().UnixNano())),
		Lock: &sync.RWMutex{},
	}

	for i := 0; i < 10000000; i++ {

		go func(o RandWithRW) {
			for i := 0; i < 100; i++ {
				o.Lock.Lock()
				n := o.R.Intn(1000)
				o.Lock.Unlock()
				l[n]++
			}
		}(r)
	}

	time.Sleep(time.Second)

	sort.Ints(l)

	fmt.Println(l[999] - l[0])

}
