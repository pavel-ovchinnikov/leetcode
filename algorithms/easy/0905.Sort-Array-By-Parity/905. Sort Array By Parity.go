package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	n := len(nums)
	l, r := 0, n-1

	for l < r {
		if nums[l]%2 != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			r--
		} else {
			l++
		}
	}

	return nums
}

func main() {
	fmt.Println(sortArrayByParity([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
