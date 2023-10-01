package main

import "fmt"

func dominantIndex(nums []int) int {
	maxIdx, prevIdx := 0, 0

	if nums[0] > nums[1] {
		maxIdx, prevIdx = 0, 1
	} else {
		maxIdx, prevIdx = 1, 0
	}

	for i := 2; i < len(nums); i++ {
		if nums[maxIdx] <= nums[i] {
			prevIdx, maxIdx = maxIdx, i
			continue
		}

		if nums[prevIdx] < nums[i] {
			prevIdx = i
		}
	}

	if nums[maxIdx] >= nums[prevIdx]*2 {
		return maxIdx
	}

	return -1
}

func main() {
	fmt.Println(dominantIndex([]int{3, 6, 1, 0}))
}
