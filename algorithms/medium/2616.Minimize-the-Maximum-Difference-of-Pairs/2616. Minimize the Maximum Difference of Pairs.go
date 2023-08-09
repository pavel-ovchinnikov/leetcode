package main

import (
	"fmt"
	"sort"
)

func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	l, r := 0, nums[len(nums)-1]-nums[0]
	res := r
	for l <= r {
		mid := l + (r-l)/2
		if check(nums, p, mid) {
			if mid < res {
				res = mid
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}

func check(nums []int, p, mid int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] <= mid {
			p--
			i++
		}
	}
	return p <= 0
}

func main() {
	fmt.Println(minimizeMax([]int{10, 1, 2, 7, 1, 3}, 2))
}
