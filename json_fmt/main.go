package main

import "fmt"

func main() {

	s := ""
	buf := make([]byte, 3)

	copy(buf[3:], s)

	fmt.Println(string(buf))

}
