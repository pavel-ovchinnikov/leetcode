package main

import "fmt"

func islandPerimeter(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	lands, neighbours := 0, 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] != 1 {
				continue
			}

			lands++

			if i < n-1 && grid[i+1][j] == 1 {
				neighbours++
			}

			if j < m-1 && grid[i][j+1] == 1 {
				neighbours++
			}
		}
	}
	return lands*4 - neighbours*2
}

func main() {
	fmt.Println(islandPerimeter([][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}))
}
