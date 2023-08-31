package main

import (
	"fmt"
	"sort"
)

func minimumCost(cost []int) int {
	n := len(cost)
	res := 0
	sort.Slice(cost, func(i, j int) bool {
		return cost[i] > cost[j]
	})

	for i := 0; i < n; i++ {
		if i%3 == 2 {
			continue
		}
		res += cost[i]
	}

	return res
}

func main() {
	fmt.Println(minimumCost([]int{6, 5, 7, 9, 2, 2}))
	fmt.Println(minimumCost([]int{1, 2, 3, 3, 3}))
}
