package main

import (
	"fmt"
	"sort"
)

func maxCoins(piles []int) int {
	sort.Ints(piles)
	res := 0
	n := len(piles)
	for i := n / 3; i < n; i += 2 {
		res += piles[i]
	}
	return res
}

func main() {
	fmt.Println(maxCoins([]int{9, 8, 7, 6, 5, 1, 2, 3, 4}))
}
