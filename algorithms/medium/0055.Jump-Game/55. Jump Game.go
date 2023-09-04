package main

import "fmt"

func canJump(nums []int) bool {
	n := len(nums)
	r := 0
	for i := 0; i < n && i <= r; i++ {
		r = max(i+nums[i], r)
	}
	return r >= n-1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// func canJump(nums []int) bool {
// 	n := len(nums)
// 	l, r := 0, nums[0]
// 	for i := 0; i < n; i++ {
// 		if i > r {
// 			return false
// 		}

// 		if i+nums[i] > r {
// 			l = i
// 			r = l + nums[l]
// 		}

// 	}
// 	return r >= n-1
// }

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}))
}
