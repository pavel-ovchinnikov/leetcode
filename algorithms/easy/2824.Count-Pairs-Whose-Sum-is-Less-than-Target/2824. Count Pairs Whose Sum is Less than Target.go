package main

import (
	"fmt"
	"sort"
)

func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	res := 0
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l]+nums[r] < target {
			res += r - l
			l++
		} else {
			r--
		}
	}
	return res
}

func main() {
	fmt.Println(countPairs([]int{-1, 1, 2, 3, 1}, 2))
}
