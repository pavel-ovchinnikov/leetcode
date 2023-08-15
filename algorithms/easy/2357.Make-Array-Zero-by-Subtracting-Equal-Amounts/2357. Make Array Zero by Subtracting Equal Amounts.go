package main

import (
	"fmt"
	"sort"
)

func minimumOperations(nums []int) int {
	sort.Ints(nums)
	count := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			continue
		}

		if nums[i]-sum < 0 {
			continue
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		sum += (nums[i] - sum)
		count++
	}
	return count
}

func main() {
	fmt.Println(minimumOperations([]int{1, 5, 0, 3, 5}))
}
