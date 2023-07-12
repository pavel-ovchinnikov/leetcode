package main

import "fmt"

func moveZeroes(nums []int) {
	j := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
		}
		if nums[j] != 0 {
			j++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println()
}
