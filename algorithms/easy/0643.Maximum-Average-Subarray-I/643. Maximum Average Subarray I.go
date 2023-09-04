package main

import (
	"fmt"
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	maxSum := sum(nums[:k])
	prevSum := maxSum

	for i := 1; i <= len(nums)-k; i++ {
		newSum := prevSum - float64(nums[i-1]) + float64(nums[i+k-1])
		prevSum = newSum
		maxSum = math.Max(maxSum, newSum)
	}

	return maxSum / float64(k)
}

func sum(nums []int) (sum float64) {
	for _, num := range nums {
		sum += float64(num)
	}
	return
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
}
