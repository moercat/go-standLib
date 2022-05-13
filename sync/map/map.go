package main

import (
	"fmt"
	"sync"
)

//
//type Map struct {
//	// 该锁用来保护dirty
//	mu Mutex
//	// 存读的数据，因为是atomic.value类型，只读类型，所以它的读是并发安全的
//	read atomic.Value // readOnly
//	//包含最新的写入的数据，并且在写的时候，会把read 中未被删除的数据拷贝到该dirty中，因为是普通的map存在并发安全问题，需要用到上面的mu字段。
//	dirty map[interface{}]*entry
//	// 从read读数据的时候，会将该字段+1，当等于len（dirty）的时候，会将dirty拷贝到read中（从而提升读的性能）。
//	misses int
//}

func main() {
	var sm sync.Map

	sm.Store(1, "a") //store 方法,添加元素

	if v, ok := sm.Load(1); ok { //Load 方法，获得value
		fmt.Println(v)
	}
	//LoadOrStore方法，获取或者保存
	//参数是一对key：value，如果该key存在且没有被标记删除则返回原先的value（不更新）和true；不存在则store，返回该value 和false
	if vv, ok := sm.LoadOrStore(1, "c"); ok {
		fmt.Println(vv)
	}
	if vv, ok := sm.LoadOrStore(2, "c"); !ok {
		fmt.Println(vv)
	}
	//遍历该map，参数是个函数，该函数参的两个参数是遍历获得的key和value，返回一个bool值，当返回false时，遍历立刻结束。
	sm.Range(func(k, v interface{}) bool {
		fmt.Printf("%v:%v \n", k, v)
		return true
	})

	//LoadAndDelete，获取后删除
	//参数是key，如果该key存在则返回原先的value  和 true 且删除；不存在则返回 nil 和 false
	if vv, ok := sm.LoadAndDelete(3); !ok {
		fmt.Println(vv, ok)
	}
	sm.Range(func(k, v interface{}) bool {
		fmt.Printf("%v:%v \n", k, v)
		return true
	})
	//Delete，删除键的值。
	sm.Delete(1)

	sm.Range(func(k, v interface{}) bool {
		fmt.Printf("%v:%v \n", k, v)
		return true
	})

}
