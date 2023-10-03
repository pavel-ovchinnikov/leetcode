package main

import (
	"fmt"
	"sort"
)

func numIdenticalPairs(nums []int) int {
	cnt := make(map[int]int)
	var pairs int

	for _, num := range nums {
		pairs += cnt[num]
		cnt[num]++
	}
	return pairs
}

func numIdenticalPairs2(nums []int) int {
	n := len(nums)
	count := 1
	res := 0
	sort.Ints(nums)

	for i := 1; i < n; i++ {
		if nums[i-1] == nums[i] {
			count++
		} else {
			res += count * (count - 1) / 2
			count = 1
		}
	}
	res += count * (count - 1) / 2

	return res
}

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
	fmt.Println(numIdenticalPairs2([]int{1, 2, 3, 1, 1, 3}))
}
