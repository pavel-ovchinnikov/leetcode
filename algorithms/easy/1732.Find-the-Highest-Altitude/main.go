package main

import "fmt"

func largestAltitude(gain []int) int {
	maxVal := 0
	cur := 0
	for _, val := range gain {
		cur += val
		if cur > maxVal {
			maxVal = cur
		}
	}
	return maxVal
}

func main() {
	fmt.Println(largestAltitude([]int{-5, 1, 5, 0, -7}))
	fmt.Println(largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
}
