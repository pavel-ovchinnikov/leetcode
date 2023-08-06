package main

import "fmt"

func minOperations(nums []int) int {
	res := 0

	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			diff := nums[i-1] + 1 - nums[i]
			nums[i] += diff
			res += diff
		}
	}
	return res
}

func main() {
	fmt.Println(minOperations([]int{1, 5, 2, 4, 1}))
}
