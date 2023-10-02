package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	res := 0
	l, r := 0, 0
	n := len(nums)
	for l < n {
		r = l + 1
		for r < n && nums[r] > nums[r-1] {
			r++
		}
		res = Max(res, r-l)
		l = r
	}
	return res
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 2, 3, 1, 2}))
}
