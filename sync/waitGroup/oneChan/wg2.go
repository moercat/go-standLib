package main

import (
	"fmt"
	"sync"
)

// 两个线程交替打印0~100的奇偶数
var wg sync.WaitGroup

func j(ch chan struct{}) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch <- struct{}{}
		if i%2 == 1 {
			fmt.Println(i)
		}
	}
}

func o(ch chan struct{}) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		<-ch
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	wg.Add(2)
	ch := make(chan struct{})

	go j(ch)
	go o(ch)

	wg.Wait()
}
