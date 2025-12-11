package quicksort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "empty slice",
			arr:  []int{},
			want: []int{},
		},
		{
			name: "single element",
			arr:  []int{5},
			want: []int{5},
		},
		{
			name: "two elements sorted",
			arr:  []int{1, 2},
			want: []int{1, 2},
		},
		{
			name: "two elements unsorted",
			arr:  []int{2, 1},
			want: []int{1, 2},
		},
		{
			name: "multiple elements with duplicates",
			arr:  []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
			want: []int{1, 1, 2, 3, 4, 5, 5, 6, 9},
		},
		{
			name: "already sorted array",
			arr:  []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "reverse sorted array",
			arr:  []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "array with negative numbers",
			arr:  []int{-3, 1, -4, 0, 5},
			want: []int{-4, -3, 0, 1, 5},
		},
		{
			name: "array with all same elements",
			arr:  []int{7, 7, 7, 7},
			want: []int{7, 7, 7, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeSort(tt.arr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
