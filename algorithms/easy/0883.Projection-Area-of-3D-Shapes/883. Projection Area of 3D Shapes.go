package main

import "fmt"

func projectionArea(grid [][]int) int {
	res := 0
	n := len(grid)
	x, y, z := 0, 0, 0

	for i := 0; i < n; i++ {
		x, y, z = 0, 0, 0
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				x += 1
			}
			y = max(y, grid[i][j])
			z = max(z, grid[j][i])
		}
		res += x + y + z
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
	// fmt.Println(projectionArea(
	// 	[][]int{
	// 		{1, 2},
	// 		{3, 4}}))

	fmt.Println(projectionArea(
		[][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1}}))
}
