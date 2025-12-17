package collections

import "fmt"

// CollectionsDemo 运行collections包中的所有演示功能
func CollectionsDemo() {
	fmt.Println("=== Running List Demo ===")
	ListDemo()

	fmt.Println("\n=== Running Ring Demo ===")
	RingDemo()

	fmt.Println("\n=== Running Slice Range Demo ===")
	SliceRangeDemo()

	fmt.Println("\n=== Running Slice Window Demo ===")
	SliceWindowDemo()
}
