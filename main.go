package main

import (
	"fmt"
	"go-standLib/quicksort"
)

func main() {
	// 示例数组
	arr := []int{64, 34, 25, 12, 22, 11, 90, 5}
	fmt.Println("Original array:", arr)

	// 执行快速排序
	quicksort.QuickSort(arr)

	fmt.Println("Sorted array:  ", arr)

	// 另一个示例
	arr2 := []int{3, 6, 8, 10, 1, 2, 1}
	fmt.Println("\nOriginal array:", arr2)

	quicksort.QuickSort(arr2)

	fmt.Println("Sorted array:  ", arr2)
}
