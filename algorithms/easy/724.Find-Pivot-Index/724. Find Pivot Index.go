package main

import "fmt"

func pivotIndex(nums []int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	prefSum, postSum := 0, sum

	for i := 0; i < n; i++ {
		postSum -= nums[i]
		if prefSum == postSum {
			return i
		}

		prefSum += nums[i]
	}
	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}
