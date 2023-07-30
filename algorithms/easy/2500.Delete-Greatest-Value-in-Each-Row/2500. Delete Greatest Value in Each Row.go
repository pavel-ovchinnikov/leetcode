package main

import (
	"fmt"
	"sort"
)

func deleteGreatestValue(grid [][]int) int {
	res := 0

	for i := 0; i < len(grid); i++ {
		sort.Ints(grid[i])
	}

	for i := 0; i < len(grid[0]); i++ {
		tmp := 0
		for j := 0; j < len(grid); j++ {
			tmp = max(tmp, grid[j][i])
		}
		res += tmp
	}
	return res
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(deleteGreatestValue([][]int{
		{1, 2, 4},
		{3, 3, 1},
	}))
}
