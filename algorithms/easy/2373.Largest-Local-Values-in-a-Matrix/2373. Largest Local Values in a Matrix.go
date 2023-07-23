package main

import "fmt"

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	result := make([][]int, n-2)
	for index := range result {
		result[index] = make([]int, n-2)
	}

	for i := 0; i < n-2; i++ {
		for j := 0; j < n-2; j++ {
			result[i][j] = max(
				grid[i][j], grid[i][j+1], grid[i][j+2],
				grid[i+1][j], grid[i+1][j+1], grid[i+1][j+2],
				grid[i+2][j], grid[i+2][j+1], grid[i+2][j+2])
		}
	}
	return result
}

func max(nums ...int) int {
	var max = 0
	for _, number := range nums {
		if max < number {
			max = number
		}
	}
	return max
}

func main() {
	fmt.Println(largestLocal([][]int{
		{9, 9, 8, 1},
		{5, 6, 2, 6},
		{8, 2, 6, 4},
		{6, 2, 2, 2},
	}))
}
