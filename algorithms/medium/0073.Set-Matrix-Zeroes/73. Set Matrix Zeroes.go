package main

import "fmt"

func setZeroes(matrix [][]int) {
	points := [][]int{}
	n := len(matrix)
	m := len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				points = append(points, []int{i, j})
			}
		}
	}

	for _, point := range points {
		x, y := point[0], point[1]
		for i := 0; i < n; i++ {
			matrix[i][y] = 0
		}
		for j := 0; j < m; j++ {
			matrix[x][j] = 0
		}
	}
}

func main() {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(matrix)
	for _, row := range matrix {
		fmt.Println(row)
	}

}
