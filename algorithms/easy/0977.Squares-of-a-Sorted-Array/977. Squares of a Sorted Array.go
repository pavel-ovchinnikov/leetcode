package main

import "fmt"

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	i := n - 1
	l, r := 0, n-1
	for l <= r {
		if abs(nums[l]) >= abs(nums[r]) {
			res[i] = nums[l] * nums[l]
			l++
		} else {
			res[i] = nums[r] * nums[r]
			r--
		}
		i--
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
}
