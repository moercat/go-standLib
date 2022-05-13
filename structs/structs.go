package main

import "fmt"

type Test struct {
	age int
	_   struct{}
}

func main() {
	a := Test{
		age: 1,
	}

	b := a

	fmt.Println(a == b)
}
