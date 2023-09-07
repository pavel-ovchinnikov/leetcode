package main

import "fmt"

func findMin(nums []int) int {
	n := len(nums)
	l, r, m := 0, n-1, 0
	for l < r {
		m = l + (r-l)/2
		if nums[m] <= nums[r] {
			r = m
		} else {
			l = m + 1
		}
	}

	return nums[l]
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{0, 1, 2, 4, 5, 6, 7}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
	fmt.Println(findMin([]int{11, 13, 15, 17, 10}))
	fmt.Println(findMin([]int{100, 11, 13, 15, 17}))
}
