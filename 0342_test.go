package kojiro

import "testing"

func TestIsPowerOf4(t *testing.T) {
	tests := []struct {
		testCase   int
		testResult bool
	}{
		{testCase: 16, testResult: true},
		{testCase: 5, testResult: false},
		{testCase: 1, testResult: true},
		{testCase: -1, testResult: false},
	}

	for testIdx := 0; testIdx < len(tests); testIdx++ {
		test := tests[testIdx]
		if test.testResult != isPowerOfFour(test.testCase) {
			t.Errorf("Test failed for %d", test.testCase)
		}
	}
}

func TestBitwiseIsPowerOf4(t *testing.T) {
	tests := []struct {
		testCase   int
		testResult bool
	}{
		{testCase: 16, testResult: true},
		{testCase: 5, testResult: false},
		{testCase: 1, testResult: true},
		{testCase: -1, testResult: false},
	}

	for testIdx := 0; testIdx < len(tests); testIdx++ {
		test := tests[testIdx]
		if test.testResult != bitwiseIsPowerOfFour(test.testCase) {
			t.Errorf("Test failed for %d", test.testCase)
		}
	}
}
