// 消息队列
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	catEnd := make(chan struct{}, 1)
	dogEnd := make(chan struct{}, 1)
	fishEnd := make(chan struct{}, 1)
	fishEnd <- struct{}{}
	go Cat(fishEnd, catEnd)
	go Dog(catEnd, dogEnd)
	go Fish(dogEnd, fishEnd)

	wg.Wait()
}

func Cat(fishEnd chan struct{}, catEnd chan struct{}) {
	for i := 0; i < 100; i++ {
		<-fishEnd
		fmt.Println("cat", i)
		catEnd <- struct{}{}
	}
	wg.Done()
}

func Dog(catEnd chan struct{}, dogEnd chan struct{}) {
	for i := 0; i < 100; i++ {
		<-catEnd
		fmt.Println("dog", i)
		dogEnd <- struct{}{}
	}
	wg.Done()
}

func Fish(dogEnd chan struct{}, fishEnd chan struct{}) {
	for i := 0; i < 100; i++ {
		<-dogEnd
		fmt.Println("fish", i)
		fishEnd <- struct{}{}
	}
	wg.Done()
}
