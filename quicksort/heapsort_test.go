package quicksort

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
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
			name:     "Two elements",
			input:    []int{2, 1},
			expected: []int{1, 2},
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
			name:     "With duplicates",
			input:    []int{4, 2, 7, 2, 4, 8, 1},
			expected: []int{1, 2, 2, 4, 4, 7, 8},
		},
		{
			name:     "With negative numbers",
			input:    []int{-3, 6, -8, 10, 1, -2, 0},
			expected: []int{-8, -3, -2, 0, 1, 6, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of the input array to avoid modifying the original test data
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			HeapSort(inputCopy)

			if !reflect.DeepEqual(inputCopy, tt.expected) {
				t.Errorf("HeapSort() = %v, want %v", inputCopy, tt.expected)
			}
		})
	}
}

func TestHeapSortInPlace(t *testing.T) {
	// Test that the sort works in-place and modifies the original array
	arr := []int{5, 2, 8, 1, 9}
	expected := []int{1, 2, 5, 8, 9}

	HeapSort(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("HeapSort() modified array = %v, want %v", arr, expected)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{3, 6, 8, 10, 1, 2, 1, 7, 9, 5, 4, 11, 13, 12, 15, 14}
		HeapSort(arr)
	}
}

func BenchmarkHeapSortLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]int, 1000)
		for j := 0; j < 1000; j++ {
			arr[j] = 1000 - j
		}
		HeapSort(arr)
	}
}
