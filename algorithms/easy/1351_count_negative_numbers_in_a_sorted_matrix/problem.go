package problem

func countNegatives(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		row := grid[i]
		for j := 0; j < len(row); j++ {
			if grid[i][j] < 0 {
				count++
			}
		}
	}
	return count
}
