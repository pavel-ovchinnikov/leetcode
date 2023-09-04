package main

import (
	"fmt"
	"math"
)

//// Variant 1
// func maxSubArray(nums []int) int {
// 	n := len(nums)
// 	maxSum := nums[0]
// 	for i := 0; i < n; i++ {
// 		sum := 0
// 		for j := i; j < n; j++ {
// 			sum += nums[j]
// 			maxSum = max(maxSum, sum)
// 		}
// 	}
// 	return maxSum
// }

//// Variant 2
// func maxSubArray(nums []int) int {
// 	n := len(nums)
// 	dp := make([]int, n)
// 	for i := 1; i < len(nums); i++ {
// 		dp[i] = max(nums[i], nums[i]+dp[i-1])
// 	}

// 	return max(dp...)
// }

// Variant 3
func maxSubArray(nums []int) int {
	curMax, maxSub := 0, math.MinInt
	for _, num := range nums {
		curMax = max(curMax+num, num)
		maxSub = max(maxSub, curMax)
	}
	return maxSub
}

func max(nums ...int) int {
	maxVal := math.MinInt
	for _, num := range nums {
		if maxVal < num {
			maxVal = num
		}
	}
	return maxVal
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
