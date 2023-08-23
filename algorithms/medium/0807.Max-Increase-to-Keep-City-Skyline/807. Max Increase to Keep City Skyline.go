package main

import "fmt"

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	west := make([]int, n)
	north := make([]int, m)
	res := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			west[i] = max(west[i], grid[i][j])
			north[j] = max(north[j], grid[i][j])
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res += min(west[i], north[j]) - grid[i][j]
		}
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxIncreaseKeepingSkyline([][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0},
	}))
}
