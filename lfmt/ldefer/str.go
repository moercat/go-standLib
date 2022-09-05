package main

import "fmt"

func main() {
	fmt.Println("end i", deferT())
}

func deferT() (s int) {
	var i = 5

	defer fmt.Println("1i", i)

	defer func() {
		fmt.Println("2i", i)
	}()

	defer func(j int) {
		fmt.Println("3i", j)
	}(i)

	defer func() {
		i += 10
		fmt.Println("4i", i)
	}()

	defer func() {
		s += 10
		fmt.Println("5i", s)
	}()

	defer func(s int) {
		s = s + 1
		fmt.Println("6i", s)
	}(s)

	i += 10

	fmt.Println("begin i", i)

	return i
}
