package kojiro

import "math"

func maximum69Number(num int) int {
	numberArrayReversed := []int{}
	copyOfNum := num

	for num > 0 {
		numberArrayReversed = append(numberArrayReversed, num%10)
		num /= 10
	}

	last6Idx := -1
	for idx := 0; idx < len(numberArrayReversed); idx++ {
		if numberArrayReversed[idx] == 6 {
			last6Idx = idx
		}
	}

	if last6Idx == -1 {
		return copyOfNum
	}

	return copyOfNum + 3*int(math.Pow10(last6Idx))
}
