package main

import (
	"fmt"
)

type node struct {
	x, y, dist int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	l := len(grid)
	n := l - 1
	if grid[0][0] == 1 || grid[n][n] == 1 {
		return -1
	}
	queue := make([]node, 0)
	queue = append(queue, node{x: 0, y: 0, dist: 1})
	yarr := [8]int{1, 1, 0, -1, -1, -1, 0, 1}
	xarr := [8]int{0, 1, 1, 1, 0, -1, -1, -1}
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		if front.x == n && front.y == n {
			return front.dist
		}
		for k := 0; k < len(xarr); k++ {
			i, j := xarr[k]+front.x, yarr[k]+front.y
			if i >= 0 && i < l && j >= 0 && j < l && grid[i][j] == 0 {
				queue = append(queue, node{x: i, y: j, dist: front.dist + 1})
				grid[i][j] = 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(shortestPathBinaryMatrix([][]int{
		{0, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}))
	fmt.Println(shortestPathBinaryMatrix([][]int{
		{1, 0, 0},
		{1, 1, 0},
		{1, 1, 0},
	}))
	fmt.Println(shortestPathBinaryMatrix([][]int{
		{0},
	}))

	fmt.Println(shortestPathBinaryMatrix([][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 1, 0},
		{0, 0, 1, 0, 1, 0, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 0, 0},
		{1, 0, 1, 1, 1, 0, 0, 0},
	}))

}
