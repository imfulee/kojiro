package main

func getPenalty(hours string, closed int) int {
	if closed > len(hours) {
		closed = len(hours)
	}

	penalty := 0

	for idx := 0; idx < closed; idx++ {
		if hours[idx] == 'N' {
			penalty++
		}
	}
	for idx := closed; idx < len(hours); idx++ {
		if hours[idx] == 'Y' {
			penalty++
		}
	}

	return penalty
}

func bestClosingTime(customers string) int {
	smallestPenalty := getPenalty(customers, len(customers))
	closingHour := len(customers)
	lastPenalty := smallestPenalty

	for hour := len(customers) - 1; hour >= 0; hour-- {
		p := lastPenalty
		switch customers[hour] {
		case 'Y':
			p++
		case 'N':
			p--
		}
		lastPenalty = p

		if p <= smallestPenalty {
			closingHour = hour
			smallestPenalty = p
		}
	}

	return closingHour
}
