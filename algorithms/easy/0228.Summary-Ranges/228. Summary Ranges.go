package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	n := len(nums)
	if n == 0 {
		return []string{}
	}

	res := make([]string, 0)

	l := 0
	r := 1

	for r < n {
		if nums[r]-nums[r-1] == 1 {
			r++
		} else {
			res = append(res, covert(nums[l], nums[r-1]))
			l = r
			r = l + 1
		}
	}

	if l != r {
		res = append(res, covert(nums[l], nums[n-1]))
	}

	return res
}

func covert(a, b int) string {
	if a == b {
		return strconv.FormatInt(int64(a), 10)
	}
	start := strconv.FormatInt(int64(a), 10)
	finish := strconv.FormatInt(int64(b), 10)
	return start + "->" + finish
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{1, 4}))
	fmt.Println(summaryRanges([]int{}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	fmt.Println(summaryRanges([]int{-1, 2, 3, 4, 6, 8, 11, 12}))
}
