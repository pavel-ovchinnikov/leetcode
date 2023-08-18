package main

import (
	"fmt"
	"sort"
)

func frequencySort(nums []int) []int {
	hash := make(map[int]int)
	for _, v := range nums {
		hash[v]++
	}

	sort.Slice(nums, func(i, j int) bool {
		if hash[nums[i]] == hash[nums[j]] {
			return nums[i] > nums[j]
		}
		return hash[nums[i]] < hash[nums[j]]
	})
	return nums
}

func main() {
	fmt.Println(frequencySort([]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}))
}
