package kojiro

func isPowerOfFour(n int) bool {
	floatN := float64(n)

	for floatN >= 4 {
		floatN /= 4
	}

	return floatN == 1
}

func bitwiseIsPowerOfFour(n int) bool {
	return n > 0 && (n&(n-1)) == 0 && n%3 == 1
}
