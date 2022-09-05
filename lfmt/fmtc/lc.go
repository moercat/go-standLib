package main

import "fmt"

func main() {
	lc()
}

func lc() {
	for i := 0; i < 25; i++ {
		if i < 12 {
			for j := 0; j < (12 - i); j++ {
				fmt.Printf("-")
			}
			fmt.Println("*")
		} else {
			for j := 0; j < (i - 12); j++ {
				fmt.Printf("-")
			}
			fmt.Println("*")
		}

	}

}
