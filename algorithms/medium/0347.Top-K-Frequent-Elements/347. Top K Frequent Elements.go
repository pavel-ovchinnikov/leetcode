package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	counter := map[int]int{}
	for _, num := range nums {
		counter[num]++
	}

	res := make([]int, 0, len(counter))
	for key := range counter {
		res = append(res, key)
	}
	sort.Slice(res, func(i, j int) bool {
		return counter[res[i]] > counter[res[j]]
	})

	return res[:k]
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3, 4}, 2))
}
