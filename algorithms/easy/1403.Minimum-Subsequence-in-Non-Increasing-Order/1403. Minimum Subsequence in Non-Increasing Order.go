package main

import (
	"fmt"
	"sort"
)

func minSubsequence(nums []int) []int {
	res := []int{}
	sum := 0
	resSum := 0

	sort.Ints(nums)

	for _, num := range nums {
		sum += num
	}

	for i := len(nums) - 1; i >= 0 && resSum <= sum; i-- {
		sum -= nums[i]
		resSum += nums[i]
		res = append(res, nums[i])
	}
	return res
}

func main() {
	fmt.Println(minSubsequence([]int{4, 3, 10, 9, 8}))
}
