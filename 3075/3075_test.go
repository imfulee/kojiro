package main

import (
	"testing"
)

func TestMaxHappinessSumCase1(t *testing.T) {
	sum := maximumHappinessSum([]int{1, 2, 3}, 2)
	if sum != 4 {
		t.Errorf("Happiness sum for 1,2,3 | k=2 should be 4, but is %d", sum)
	}
}

func TestMaxHappinessSumCase2(t *testing.T) {
	sum := maximumHappinessSum([]int{1, 1, 1, 1}, 2)
	if sum != 1 {
		t.Errorf("Happiness sum for 1,1,1,1 | k=2 should be 1, but is %d", sum)
	}
}

func TestMaxHappinessSumCase3(t *testing.T) {
	sum := maximumHappinessSum([]int{2, 3, 4, 5}, 1)
	if sum != 5 {
		t.Errorf("Happiness sum for 2,3,4,5 | k=1 should be 5, but is %d", sum)
	}
}
