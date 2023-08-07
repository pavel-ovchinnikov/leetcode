package main

import (
	"fmt"
	"math"
)

func findGCD(nums []int) int {
	maxVal, minVal := math.MinInt, math.MaxInt

	for _, val := range nums {
		maxVal = max(maxVal, val)
		minVal = min(minVal, val)
	}
	return GCD(minVal, maxVal)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func GCD(a, b int) int {
	for a > 0 {
		a, b = b%a, a
	}
	return b
}

func main() {
	// fmt.Println(GCD(10, 2))
	fmt.Println(findGCD([]int{2, 5, 6, 9, 10}))
}
