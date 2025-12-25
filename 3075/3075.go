package main

import (
	"slices"
)

func maximumHappinessSum(happiness []int, k int) int64 {
	sorted := slices.Sorted(slices.Values(happiness))
	slices.Reverse(sorted)
	sum := 0
	for i := range k {
		h := sorted[i] - i
		if h <= 0 {
			break
		}
		sum += h
	}
	return int64(sum)
}
