package main

import "fmt"

func maxProduct(nums []int) int {
	minVal, maxVal, res := 1, 1, nums[0]
	for _, n := range nums {
		minNum, maxNum := minVal, maxVal
		minVal = min(n, min(minNum*n, maxNum*n))
		maxVal = max(n, max(minNum*n, maxNum*n))
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
