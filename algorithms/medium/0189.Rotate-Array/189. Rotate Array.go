package main

import "fmt"

// func rotate(nums []int, k int) {
// 	n := len(nums)
// 	res := make([]int, len(nums))
// 	for i, num := range nums {
// 		res[(i+k)%n] = num
// 	}
// 	copy(nums, res)
// }

func rotate(nums []int, k int) {
	n := len(nums)
	for i := 0; i < k; i++ {
		tmp := nums[n-1]
		for j := n - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = tmp
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}
