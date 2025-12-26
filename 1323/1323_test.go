package kojiro

import "testing"

func TestMaximum69Number(t *testing.T) {
	tests := []struct {
		testCase   int
		testResult int
	}{{testCase: 9669, testResult: 9969}, {testCase: 9996, testResult: 9999}, {testCase: 9999, testResult: 9999}}

	for testIdx := 0; testIdx < len(tests); testIdx++ {
		test := tests[testIdx]
		if test.testResult != maximum69Number(test.testCase) {
			t.Errorf("Test failed for %d", test.testCase)
		}
	}
}
