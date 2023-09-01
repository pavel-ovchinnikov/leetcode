package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for l := 0; l < len(nums); l++ {
		if l > 0 && nums[l] == nums[l-1] {
			continue
		}

		for r := len(nums) - 1; r > l+1; r-- {
			if r < len(nums)-1 && nums[r+1] == nums[r] {
				continue
			}

			second := -(nums[l] + nums[r])
			if second > nums[r] {
				break
			}
			if binSearch(nums, l+1, r-1, second) != -1 {
				res = append(res, []int{nums[l], second, nums[r]})
			}
		}
	}
	return res
}

func binSearch(a []int, l, r, t int) int {
	for l <= r {
		m := (l + r) / 2

		if a[m] == t {
			return m
		}

		if t > a[m] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}

func main() {
	// fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	// fmt.Println(threeSum([]int{-1, 0, 1, 2}))
	fmt.Println(threeSum([]int{0, 0, 0}))
}
