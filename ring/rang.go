package main

import (
	"container/ring"
	"fmt"
)

var r = ring.New(5)
var s = ring.New(5)

func main() {

	fmt.Println("len", r.Len())

	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()

		s.Value = i + r.Len()
		s = s.Next()
	}

	r.Move(4).Link(s)

	r.Do(func(a any) {
		fmt.Println(a)
	})

	r = r.Prev()

	r.Unlink(3)

	r.Next().Do(func(a any) {
		fmt.Println(a)
	})

}
