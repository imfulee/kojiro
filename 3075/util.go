package main

func MergeSort[E any](array []E, compare func(a E, b E) int) []E {
	newArray := []E{}

	if len(array) == 0 || len(array) == 1 {
		newArray = array
	} else if len(array) >= 2 {
		half := len(array) / 2
		left := MergeSort(array[:half], compare)
		right := MergeSort(array[half:], compare)
		leftIdx := 0
		rightIdx := 0

		for leftIdx < len(left) || rightIdx < len(right) {
			if leftIdx >= len(left) {
				newArray = append(newArray, right[rightIdx:]...)
				break
			}
			if rightIdx >= len(right) {
				newArray = append(newArray, left[leftIdx:]...)
				break
			}

			compared := compare(left[leftIdx], right[rightIdx])
			if compared < 0 {
				newArray = append(newArray, left[leftIdx])
				leftIdx++
			} else if compared > 0 {
				newArray = append(newArray, right[rightIdx])
				rightIdx++
			} else {
				newArray = append(newArray, left[leftIdx], right[rightIdx])
				leftIdx++
				rightIdx++
			}
		}
	}

	return newArray
}
