package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	if obstacleGrid[n-1][m-1] == 1 || obstacleGrid[0][0] == 1 {
		return 0
	}
	obstacleGrid[0][0] = -1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			if i < n-1 && obstacleGrid[i+1][j] != 1 {
				obstacleGrid[i+1][j] += obstacleGrid[i][j]
			}
			if j < m-1 && obstacleGrid[i][j+1] != 1 {
				obstacleGrid[i][j+1] += obstacleGrid[i][j]
			}
		}
	}
	return -1 * obstacleGrid[n-1][m-1]
}

func main() {
	grid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(grid))
}
