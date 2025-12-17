// 消息队列
package concurrent

import (
	"fmt"
	"sync"
)

func ChannelDemo() {
	var wg sync.WaitGroup
	var pg sync.WaitGroup

	consumer := func(cname string, ch chan int) {
		//可以循环 for i := range ch 来不断从 channel 接收值，直到它被关闭。
		defer pg.Done()

	loop:
		for {
			select {
			case msg, ok := <-ch:
				if !ok {
					fmt.Println("chan closed")
					break loop
				}
				fmt.Println("consumer-----------", cname, ":", msg)
			default:
			}
		}
	}

	producer := func(pname string, ch chan int) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Println("producer--", pname, ":", i)
			ch <- i
		}
	}

	//用channel来传递"产品", 不再需要自己去加锁维护一个全局的阻塞队列
	ch := make(chan int, 4)
	wg.Add(2)
	pg.Add(2)
	go producer("生产者1", ch)
	go producer("生产者2", ch)
	go consumer("消费者1", ch)
	go consumer("消费者2", ch)

	wg.Wait()
	close(ch)
	pg.Wait()
}
