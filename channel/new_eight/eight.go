package main

import (
	"container/ring"
	"fmt"
	"sync"
	"time"
)

var eight = ring.New(8)
var wg sync.WaitGroup

func main() {
	var i int
	wg.Add(2)
	go func() {
		for {
			eight.Value = i
			eight = eight.Next()

			time.Sleep(100 * time.Millisecond)
			i++
		}

	}()

	go func() {
		for {
			eight.Value = i
			eight = eight.Next()

			time.Sleep(100 * time.Millisecond)
			i++
		}

	}()

	go func() {
		for {
			var push []int
			eight.Do(func(a any) {
				if a != nil {
					push = append(push, a.(int))
				}
			})

			fmt.Println(push)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	wg.Wait()
}
