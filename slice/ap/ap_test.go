package ap

import (
	"fmt"
	"testing"
)

func BenchmarkAdd1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s []string
		add(s)
	}
}

func BenchmarkAdd2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s = make([]string, 0, 4)
		add(s)
	}
}

func Test111(t *testing.T) {
	ListCheck()
}

func TestName(t *testing.T) {
	var a = []string{"11", "22", "33"}

	fmt.Println(a[:], a[:0])

	fmt.Println(cap(a[:]), cap(a[:0]))

	var b, c = a[:], a[:0]

	b = append(b, "44", "55")
	c = append(c, "44", "55")

	fmt.Println(b, c)

	fmt.Println(cap(b), cap(c))

	b = append(b, "66", "77")
	c = append(c, "66", "77")

	fmt.Println(b, c)

	fmt.Println(cap(b), cap(c))
}
