package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
}

func main() {

	pool := sync.Pool{New: func() interface{} {
		return new(Person)
	}}

	a := pool.Get()
	fmt.Println(a)
	b := a.(*Person)
	fmt.Println(b)
	b = &Person{
		Name: "aaa",
	}
	fmt.Println(b)
	pool.Put(b)
	fmt.Println(pool.Get())
	fmt.Println(pool.Get()) //池中不足则新增
	pool.Put(b)
	b = &Person{
		Name: "aac",
	}
	pool.Put(b)
	b = &Person{
		Name: "aab",
	}
	pool.Put(b)             //池中可以存放多个
	fmt.Println(pool.Get()) //多次拿取则随机
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())

}
