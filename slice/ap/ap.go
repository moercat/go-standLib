package main

import "fmt"

func main() {

	temp := make(map[string]int)
	fmt.Println("len(temp):", len(temp), "temp:", temp)

	temp1 := make(map[string]int, 3)
	fmt.Println("len(temp1):", len(temp1), "temp1", temp1)

}
