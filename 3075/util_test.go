package main

import (
	"slices"
	"testing"
)

func TestMergeSort(t *testing.T) {
	sorted := MergeSort([]int{1, 4, 3, 2}, func(a, b int) int {
		return a - b
	})
	if !slices.Equal(sorted, []int{1, 2, 3, 4}) {
		t.Errorf("Merge sort did not work for 1,4,3,2")
	}
}
