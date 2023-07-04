package main

import (
	"fmt"
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	n := len(nums)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}

	fmt.Println(nums, queries)
	res := make([]int, len(queries))
	for idx, query := range queries {
		res[idx] = upperBound(nums, query)
	}
	return res
}

func upperBound(array []int, target int) int {
	low, high, mid := 0, len(array)-1, 0
	for low <= high {
		mid = (low + high) / 2
		if array[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

func main() {
	fmt.Println(upperBound([]int{1, 3, 7, 12}, 3))
	fmt.Println(answerQueries([]int{4, 5, 2, 1}, []int{3, 10, 21}))
}
