package quicksort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "Already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Random order",
			input:    []int{3, 6, 8, 10, 1, 2, 1},
			expected: []int{1, 1, 2, 3, 6, 8, 10},
		},
		{
			name:     "Array with duplicates",
			input:    []int{3, 3, 3, 3},
			expected: []int{3, 3, 3, 3},
		},
		{
			name:     "Negative numbers",
			input:    []int{-3, -1, -4, -1, -5, -9, -2, -6},
			expected: []int{-9, -6, -5, -4, -3, -2, -1, -1},
		},
		{
			name:     "Mixed positive and negative",
			input:    []int{-1, 2, -3, 4, 0},
			expected: []int{-3, -1, 0, 2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 复制输入数组以避免修改原始测试数据
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			QuickSort(inputCopy)

			if !reflect.DeepEqual(inputCopy, tt.expected) {
				t.Errorf("QuickSort() = %v, want %v", inputCopy, tt.expected)
			}
		})
	}
}

// 性能测试
func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{3, 6, 8, 10, 1, 2, 1, 5, 9, 7, 4}
		QuickSort(arr)
	}
}
