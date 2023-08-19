package main

import "fmt"

func shiftGrid(grid [][]int, k int) [][]int {
	n, m := len(grid), len(grid[0])
	res := make([][]int, n)

	for i := 0; i < n; i++ {
		res[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			newI := (i + (j+k)/m) % n
			newJ := (j + k) % m
			res[newI][newJ] = grid[i][j]
		}
	}

	return res
}

func main() {
	grid := [][]int{
		{3, 8, 1, 9},
		{19, 7, 2, 5},
		{4, 6, 11, 10},
		{12, 0, 21, 13},
	}
	fmt.Println(shiftGrid(grid, 4))
}
