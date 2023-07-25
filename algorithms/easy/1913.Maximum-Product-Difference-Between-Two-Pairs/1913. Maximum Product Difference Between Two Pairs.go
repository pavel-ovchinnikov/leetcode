package main

import (
	"fmt"
	"sort"
)

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	size := len(nums)
	return nums[size-1]*nums[size-2] - nums[0]*nums[1]
}

func main() {
	fmt.Println(maxProductDifference([]int{5, 6, 2, 7, 4}))
	fmt.Println(maxProductDifference([]int{4, 2, 5, 9, 7, 4, 8}))
}
