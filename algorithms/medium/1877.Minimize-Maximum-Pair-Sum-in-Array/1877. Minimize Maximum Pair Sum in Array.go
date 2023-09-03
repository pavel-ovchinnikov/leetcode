package main

import (
	"fmt"
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)

	maxVal := 0
	l, r := 0, len(nums)-1
	for l < r {
		maxVal = max(maxVal, nums[l]+nums[r])
		l++
		r--
	}
	return maxVal

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minPairSum([]int{3, 5, 4, 2, 4, 6}))
}
