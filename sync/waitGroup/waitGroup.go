package main

import (
	"fmt"
	"sync"
)

// 两个线程交替打印0~100的奇偶数
var wg sync.WaitGroup

func j(ch, ch1 chan struct{}) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		<-ch
		if i%2 == 1 {
			fmt.Println(i)
			ch1 <- struct{}{}
		} else {
			ch <- struct{}{}
		}
	}
}

func o(ch, ch1 chan struct{}) {
	defer wg.Done()
	for i := 2; i <= 100; i++ {
		<-ch1
		if i%2 == 0 {
			fmt.Println(i)
			ch <- struct{}{}
		} else {
			ch1 <- struct{}{}
		}
	}
}

func main() {
	wg.Add(2)
	ch := make(chan struct{}, 1)
	ch1 := make(chan struct{}, 1)
	ch <- struct{}{}

	go j(ch, ch1)
	go o(ch, ch1)

	wg.Wait()
}
