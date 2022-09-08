package main

import (
	"container/ring"
	"fmt"
	"github.com/spf13/cast"
)

type Queue struct {
	*ring.Ring
	tail *ring.Ring
}

func main() {

	var a = New(10)

	a.Add(8)
	fmt.Println(a)

	a.PopFront()
	fmt.Println(a)

	a.PopBack()
	fmt.Println(a)
}

func New(v any) *Queue {
	var q Queue

	r := ring.New(1)
	r.Value = v

	q.tail = r
	q.Ring = r

	return &q
}

func (q *Queue) Add(v any) {

	r := ring.New(1)
	r.Value = v

	q.Link(r)
	q.tail = q.tail.Next()
}

func (q *Queue) PopFront() any {

	r := q
	q.Prev().Unlink(1)

	return r.Value
}

func (q *Queue) PopBack() any {
	r := q.tail

	q.tail = q.tail.Prev()
	q.tail.Unlink(1)

	return r.Value
}

func (q *Queue) String() string {
	var str string

	q.Do(func(a any) {
		if len(str) > 0 {
			str += " -> "
		}

		str += cast.ToString(a)
	})
	return str
}
