package main

import "fmt"

func maxProduct(nums []int) int {
	res := nums[0]
	n := len(nums)
	maxVal, minVal := res, res

	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			maxVal, minVal = minVal, maxVal

		}
		maxVal = max(nums[i], maxVal*nums[i])
		minVal = min(nums[i], minVal*nums[i])

		res = max(res, maxVal)
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}
