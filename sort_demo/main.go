package main

import "fmt"

var m = map[string]int{
	"A": 11,
	"B": 12,
	"C": 13,
	"D": 14,
	"E": 15,
	"F": 16,
}

var l = []string{
	"A", "B", "C", "D", "E", "F",
}

func main() {

	var (
		i      = 11
		master = "A"
	)

	for {

		if i == 5 {

		}

		if i == 5 {
			for _, v := range l {
				if m[v] < i-3 {
					master = v
				}
			}
		}

		fmt.Println(master)

		i++
	}

}
