package main

import "strings"

func hasDuplicates(chars []string) bool {
	set := map[string]bool{}
	for _, char := range chars {
		_, hasChar := set[char]
		if hasChar {
			return true
		}
		set[char] = true
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	arr := strings.Split(s, "")

	longest := 0
	startIdx := 0
	endIdx := 1

	for endIdx <= len(arr) {
		if !hasDuplicates(arr[startIdx:endIdx]) {
			length := (endIdx - startIdx)
			if length > longest {
				longest = length
			}
			endIdx++
		} else {
			startIdx++
		}
	}

	return longest
}
