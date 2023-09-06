package main

import (
	"fmt"
	"sort"
)

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	l, r := 0, n-1
	res := 0
	mod := int(1e9 + 7)
	exp := make([]int, len(nums))
	exp[0] = 1

	for i := 1; i < len(nums); i++ {
		exp[i] = (2 * exp[i-1]) % mod
	}

	for l <= r {
		if nums[l]+nums[r] <= target {
			res = (res + exp[r-l]) % mod
			l++
		} else {
			r--
		}
	}

	return res
}

func main() {
	fmt.Println(numSubseq([]int{2, 3, 3, 4, 6, 7}, 12))
}
