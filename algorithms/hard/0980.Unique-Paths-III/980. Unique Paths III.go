package main

import "fmt"

var res int = 0
var empty int = 1

func uniquePathsIII(grid [][]int) int {
	res, empty = 0, 1
	x, y := initFunc(grid)
	dfs(grid, x, y, 0)
	return res
}

func initFunc(grid [][]int) (x, y int) {
	for i, row := range grid {
		for j, val := range row {
			switch val {
			case 1:
				x, y = i, j
			case 0:
				empty++
			}
		}
	}
	return x, y
}

func dfs(grid [][]int, x, y, count int) {
	isIndexValid := (x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]))
	if !isIndexValid || grid[x][y] == -1 {
		return
	}
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == -1 {
		return
	}

	if grid[x][y] == 2 {
		if empty == count {
			res++
		}
		return
	}

	grid[x][y] = -1
	dfs(grid, x+1, y, count+1)
	dfs(grid, x-1, y, count+1)
	dfs(grid, x, y+1, count+1)
	dfs(grid, x, y-1, count+1)
	grid[x][y] = 0
}

func main() {
	fmt.Println(uniquePathsIII([][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, -1},
	}))

	fmt.Println(uniquePathsIII([][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 2},
	}))
}
