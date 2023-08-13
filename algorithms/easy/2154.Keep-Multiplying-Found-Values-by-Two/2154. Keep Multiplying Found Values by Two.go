package main

import (
	"fmt"
	"sort"
)

func findFinalValue(nums []int, original int) int {
	sort.Ints(nums)
	index := sort.SearchInts(nums, original)
	for index < len(nums) && nums[index] == original {
		original *= 2
		index = sort.SearchInts(nums, original)
	}
	return original
}

func main() {
	fmt.Println(findFinalValue([]int{5, 3, 6, 1, 12}, 3))
	fmt.Println(findFinalValue([]int{1, 2, 4, 8, 16, 32, 64}, 1))
}
