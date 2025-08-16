package kojiro

import "testing"

func TestLergestGoodInteger(t *testing.T) {
	tests := []struct {
		testCase   string
		testResult string
	}{
		{testCase: "6777133339", testResult: "777"},
		{testCase: "2300019", testResult: "000"},
		{testCase: "42352338", testResult: ""},
		{testCase: "222", testResult: "222"},
	}

	for _, test := range tests {
		if test.testResult != largestGoodInteger(test.testCase) {
			t.Errorf("Test failed for %s", test.testCase)
		}
	}
}
