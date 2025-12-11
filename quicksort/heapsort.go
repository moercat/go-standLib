package quicksort

// heapSort 使用堆排序算法对数组进行排序
func heapSort(arr []int) {
	n := len(arr)

	// 从数组构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 逐个从堆中提取元素
	for i := n - 1; i > 0; i-- {
		// 将当前根节点移到末尾
		arr[0], arr[i] = arr[i], arr[0]

		// 在缩小后的堆上调用heapify
		heapify(arr, i, 0)
	}
}

// heapify 将以索引i为根的子树转换为最大堆
// n 是堆的大小
func heapify(arr []int, n, i int) {
	largest := i     // 初始化最大值为根节点
	left := 2*i + 1  // 左子节点
	right := 2*i + 2 // 右子节点

	// 如果左子节点大于根节点
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// 如果右子节点大于目前找到的最大值
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// 如果最大值不是根节点
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i] // 交换

		// 递归地对受影响的子树进行heapify
		heapify(arr, n, largest)
	}

	for i = 0; i < n; i++ {

		heapify(arr, n, i)
	}
}

// HeapSort 是使用堆排序对数组进行排序的公共函数
func HeapSort(arr []int) {
	heapSort(arr)
}
