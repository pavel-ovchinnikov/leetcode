package main

import (
	"fmt"
)

func equalPairs(grid [][]int) int {
	var n int = len(grid)
	hash := map[[200]int]int{}
	arr := [200]int{}

	for i := 0; i < n; i++ {
		copy(arr[:], grid[i])
		hash[arr]++
	}

	res := 0
	for j := 0; j < n; j++ {
		arr = [200]int{}
		for i := 0; i < n; i++ {
			arr[i] = grid[i][j]
		}
		res += hash[arr]
	}

	return res
}

func main() {
	fmt.Println(equalPairs([][]int{
		{3, 1, 2, 2},
		{1, 4, 4, 5},
		{2, 4, 2, 2},
		{2, 4, 2, 2},
	}))
}
