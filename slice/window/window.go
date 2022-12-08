package main

import "fmt"

type ListWithNumber struct {
	Number int
	Index  int
	List   [5]int
}

func (l *ListWithNumber) Append(oldRes, res int) (newRes int) {

	if l.Index++; l.Index >= 5 {
		l.Index -= 5
	}

	l.Number = l.Number - l.List[l.Index] + res
	l.List[l.Index] = res

	if l.Number == 5 {
		return 1
	}

	if l.Number < 4 {
		return 0
	}

	return oldRes
}

func main() {

	var (
		ln  ListWithNumber
		res = 0
	)

	l := []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 1, 1, 1}

	for _, v := range l {
		res = ln.Append(res, v)
		fmt.Printf("%d ", res)
	}

}
