package main

import "testing"

func TestLongestSubstringCase1(t *testing.T) {
	test := "abcabcbb"
	expect := 3
	result := lengthOfLongestSubstring(test)
	if result != expect {
		t.Errorf("Longest of %s should be length %d but is %d", test, expect, result)
	}
}

func TestLongestSubstringCase2(t *testing.T) {
	test := "bbbbb"
	expect := 1
	result := lengthOfLongestSubstring(test)
	if result != expect {
		t.Errorf("Longest of %s should be length %d but is %d", test, expect, result)
	}
}

func TestLongestSubstringCase3(t *testing.T) {
	test := "pwwkew"
	expect := 3
	result := lengthOfLongestSubstring(test)
	if result != expect {
		t.Errorf("Longest of %s should be length %d but is %d", test, expect, result)
	}
}
