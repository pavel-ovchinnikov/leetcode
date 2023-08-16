package main

import "fmt"

func smallestRangeI(nums []int, k int) int {
	minVal, maxVal := nums[0], nums[0]

	for _, num := range nums {
		minVal, maxVal = min(minVal, num), max(maxVal, num)
	}

	if maxVal-minVal > 2*k {
		return maxVal - minVal - 2*k
	}

	return 0
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

func main() {
	fmt.Println(smallestRangeI([]int{1, 3, 6}, 2))
}
