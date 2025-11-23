package flipsquaresubmatrixvertically

func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	top, bottom := x, x+k-1

	for top < bottom {
		for j := 0; j < k; j++ {
			grid[top][y+j], grid[bottom][y+j] = grid[bottom][y+j], grid[top][y+j]
		}

		top++
		bottom--
	}
	return grid
}
