package main

import "fmt"

func minCostToMoveChips(position []int) int {
	odd, even := 0, 0
	for _, val := range position {
		if val%2 == 0 {
			odd += 1
		} else {
			even += 1
		}
	}
	if odd < even {
		return odd
	}
	return even
}

func main() {
	fmt.Println(minCostToMoveChips([]int{1, 2, 3}))
}
