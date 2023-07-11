package main

import (
	"fmt"
)

func maximumGroups(grades []int) int {
	n := len(grades)
	l, r := 0, n
	maxGroups := n + 1

	for l <= r {
		mid := l + (r-l)/2
		sum := mid * (mid + 1) / 2
		if sum <= n {
			maxGroups = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return maxGroups
}

func main() {
	fmt.Println(maximumGroups([]int{10, 6, 12, 7, 3, 5}))
}
