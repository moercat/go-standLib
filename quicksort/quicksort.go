package quicksort

// QuickSort 快速排序的主函数
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSortHelper(arr, 0, len(arr)-1)
}

// quickSortHelper 快速排序的递归辅助函数
func quickSortHelper(arr []int, low, high int) {
	if low < high {
		// 获取分区点
		pivotIndex := partition(arr, low, high)

		// 对分区点左边的子数组进行排序
		quickSortHelper(arr, low, pivotIndex-1)

		// 对分区点右边的子数组进行排序
		quickSortHelper(arr, pivotIndex+1, high)
	}
}

// partition 分区函数，将数组分为小于和大于基准值的两部分
func partition(arr []int, low, high int) int {
	// 选择最后一个元素作为基准值
	pivot := arr[high]

	// 较小元素的索引
	i := low - 1

	for j := low; j < high; j++ {
		// 如果当前元素小于或等于基准值
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // 交换元素
		}
	}

	// 将基准值放到正确位置
	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1 // 返回基准值的索引
}
