package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	size := len(nums)
	prefix := make([]int, size+1)

	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	if prefix[size] < target {
		return 0
	}

	i, j := 0, 0
	res := math.MaxInt

	for j < len(prefix) {
		diff := prefix[j] - prefix[i]
		if diff < target {
			j++
		} else {
			if j-i < res {
				res = j - i
			}
			i++
		}
	}

	return res
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}
