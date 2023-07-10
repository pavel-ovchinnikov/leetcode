package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	cnt := make(map[int]int)
	var pairs int

	for _, num := range nums {
		pairs += cnt[num]
		cnt[num]++
	}
	return pairs
}

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
}
