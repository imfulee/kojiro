package kojiro

import (
	"log"
	"strconv"
)

func findGoodSubstrings(num string) []string {
	goodSubstrings := []string{}

	for idx := 0; idx < len(num) && idx+2 < len(num); idx++ {
		if num[idx] == num[idx+1] && num[idx] == num[idx+2] {
			goodSubstrings = append(goodSubstrings, string([]byte{num[idx], num[idx+1], num[idx+2]}))
		}
	}

	return goodSubstrings
}

func largestGoodInteger(num string) string {
	goodSubstrings := findGoodSubstrings(num)

	largestGoodSubstring := ""
	for _, goodSubstring := range goodSubstrings {
		if largestGoodSubstring == "" {
			largestGoodSubstring = goodSubstring
			continue
		}

		current, err := strconv.Atoi(goodSubstring)
		if err != nil {
			log.Fatalln(err)
		}
		largest, err := strconv.Atoi(largestGoodSubstring)
		if err != nil {
			log.Fatalln(err)
		}

		if current > largest {
			largestGoodSubstring = goodSubstring
		}
	}

	return largestGoodSubstring
}
