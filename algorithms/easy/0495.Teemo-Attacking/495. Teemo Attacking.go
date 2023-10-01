package main

import "fmt"

func findPoisonedDuration(timeSeries []int, duration int) int {
	count := 0
	n := len(timeSeries)

	for i := 1; i < n; i++ {
		if timeSeries[i-1]+duration >= timeSeries[i] {
			count += timeSeries[i] - timeSeries[i-1]
		} else {
			count += duration
		}
	}

	return count + duration
}

func main() {
	fmt.Println(findPoisonedDuration([]int{1, 4}, 2))
}
