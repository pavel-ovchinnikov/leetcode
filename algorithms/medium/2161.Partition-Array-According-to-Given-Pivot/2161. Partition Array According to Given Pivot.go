package main

import "fmt"

func pivotArray(nums []int, pivot int) []int {
	res := make([]int, len(nums))

	cur := 0
	pivotFreq := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < pivot {
			res[cur] = nums[i]
			cur++
		} else if nums[i] == pivot {
			pivotFreq++
		}
	}

	for pivotFreq > 0 {
		res[cur] = pivot
		cur++
		pivotFreq--
	}

	cur = len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > pivot {
			res[cur] = nums[i]
			cur--
		}
	}

	return res
}

func main() {
	fmt.Println(pivotArray([]int{9, 12, 5, 10, 14, 3, 10}, 10))
}
