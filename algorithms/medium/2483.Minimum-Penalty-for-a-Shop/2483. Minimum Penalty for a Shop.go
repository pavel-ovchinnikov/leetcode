package main

import "fmt"

func bestClosingTime(customers string) int {
	score := 0
	maxScore := 0
	bestHour := -1

	for i, c := range customers {
		if c == 'Y' {
			score++
		} else {
			score--
		}

		if score > maxScore {
			maxScore = score
			bestHour = i
		}
	}

	return bestHour + 1
}

func main() {
	fmt.Println(bestClosingTime("YYNY"))
	fmt.Println(bestClosingTime("NNNNN"))
	fmt.Println(bestClosingTime("YYYY"))
}
