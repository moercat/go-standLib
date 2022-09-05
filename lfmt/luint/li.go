package main

import (
	"fmt"
	"math"
)

func main() {
	li()
}

func li() {
	var (
		a = uint(0) //64位系统中，uint 为 uint64
		b = uint(1)

		c = int16(100)
		d = int16(256)
	)

	fmt.Println(a - b)
	fmt.Println(c - d)

	fmt.Println(math.Pow(2, 64))

}
