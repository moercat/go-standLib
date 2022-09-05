package main

import (
	"container/list"
	"fmt"
)

var l = list.New()
var s = list.New()

func main() {

	fmt.Println("len", l.Len())

	// [5]
	fmt.Println(l.PushBack(5).Value)
	// [6,5]
	fmt.Println(l.InsertBefore(6, l.Front()).Value)
	// [6,5]
	fmt.Println(l.Front().Value)

	fmt.Println(l.Back().Value)

	s.PushBackList(l)

	fmt.Println(s.Front().Value, s.Back().Value)

	s.MoveToBack(s.Front())

	fmt.Println(s.Front().Value, s.Back().Value)

	s.Remove(s.Front())

	fmt.Println(s.Front().Value)

}
