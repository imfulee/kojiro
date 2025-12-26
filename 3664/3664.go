package main

import (
	"maps"
	"slices"
)

func scoreFromSamePosition(matchedPosX map[byte]int, bothCount int) (map[byte]int, int, int) {
	score := 0
	pXKeys := slices.Collect(maps.Keys(matchedPosX))

	for _, pXKey := range pXKeys {
		count, hasKey := matchedPosX[pXKey]
		if !hasKey || count == 0 {
			continue
		}

		for _, pairedKey := range pXKeys {
			if pairedKey == pXKey {
				continue
			}
			pairedKeyCount, hasPairedKey := matchedPosX[pairedKey]
			if !hasPairedKey || pairedKeyCount == 0 {
				break
			}

			if count >= pairedKeyCount {
				score += pairedKeyCount
				count -= pairedKeyCount
				matchedPosX[pXKey] = count
				delete(matchedPosX, pairedKey)
			} else {
				score += count
				pairedKeyCount -= count
				matchedPosX[pairedKey] = pairedKeyCount
				delete(matchedPosX, pXKey)
				count = 0
			}

			if count > 0 && bothCount > 0 {
				if count > bothCount {
					score += bothCount
					count -= bothCount
					bothCount = 0
				} else if count < bothCount {
					score += count
					bothCount -= count
					count = 0
				} else {
					score += count
					bothCount = 0
					count = 0
				}
			}

			if count == 0 {
				break
			}
		}

		if count > 0 {
			matchedPosX[pXKey] = count
		} else {
			delete(matchedPosX, pXKey)
		}
	}

	return matchedPosX, score, bothCount
}

func score(cards []string, x byte) int {
	both := 0
	matchedPos0 := map[byte]int{}
	matchedPos1 := map[byte]int{}

	for _, card := range cards {
		r0 := card[0]
		r1 := card[1]

		if r0 == x && r1 == x {
			both++
		} else if r0 == x {
			matchedPos1[r1]++
		} else if r1 == x {
			matchedPos0[r0]++
		}
	}

	matchedPos0, s0, both := scoreFromSamePosition(matchedPos0, both)
	matchedPos1, s1, both := scoreFromSamePosition(matchedPos1, both)
	score := s0 + s1

	return score
}
