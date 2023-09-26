package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			res = max(res, count)
			count = 0
		}
	}
	if count != 0 {
		res = max(res, count)
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 0, 1, 1, 1, 0, 1, 1, 1, 1}))
}
