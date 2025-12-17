package collections

import "fmt"

// SliceRangeCopyAddress 拷贝地址，原切片修改影响range的切片
func SliceRangeCopyAddress() {
	var sl = []int{1, 2, 3}
	for i, v := range sl {
		if i == 0 {
			sl[0], sl[1] = 100, 200
		}

		sl[i] = 100 + v
	}
	fmt.Println(sl)
}

// ArrayRangeCopyValue 拷贝值，原数组修改不影响range的数组
func ArrayRangeCopyValue() {
	var sl = [3]int{1, 2, 3}
	for i, v := range sl {
		if i == 0 {
			sl[0], sl[1] = 100, 200
		}

		sl[i] = 100 + v
	}
	fmt.Println(sl)
}

// SliceRangeDemo 演示切片range功能
func SliceRangeDemo() {
	fmt.Println("Slice range with address copy:")
	SliceRangeCopyAddress()

	fmt.Println("Array range with value copy:")
	ArrayRangeCopyValue()
}
