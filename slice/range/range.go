package main

import "fmt"

func main() {
	RSlice()
	RArray()
}

// RSlice 拷贝地址，原切片修改影响range的切片
func RSlice() {
	var sl = []int{1, 2, 3}
	for i, v := range sl {
		if i == 0 {
			sl[0], sl[1] = 100, 200
		}

		sl[i] = 100 + v
	}
	fmt.Println(sl)
}

// RArray 拷贝值，原数组修改不影响range的数组
func RArray() {
	var sl = [3]int{1, 2, 3}
	for i, v := range sl {
		if i == 0 {
			sl[0], sl[1] = 100, 200
		}

		sl[i] = 100 + v
	}
	fmt.Println(sl)
}
