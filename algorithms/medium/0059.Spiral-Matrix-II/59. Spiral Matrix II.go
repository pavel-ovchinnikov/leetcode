package main

import (
	"fmt"
)

const (
	Left int = iota
	Right
	Up
	Down
)

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	top, bottom := 0, n-1
	left, right := 0, n-1
	index := 1
	direction := Right

	for index <= n*n {
		switch direction {
		case Right:
			for i := left; i <= right; i++ {
				matrix[top][i] = index
				index++
			}
			top++
			direction = Down
		case Down:
			for i := top; i <= bottom; i++ {
				matrix[i][right] = index
				index++
			}
			right--
			direction = Left
		case Left:
			for i := right; i >= left; i-- {
				matrix[bottom][i] = index
				index++
			}
			bottom--
			direction = Up
		case Up:
			for i := bottom; i >= top; i-- {
				matrix[i][left] = index
				index++
			}
			left++
			direction = Right
		}
	}

	return matrix
}

func main() {
	fmt.Println(generateMatrix(1))
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(9))
	fmt.Println(generateMatrix(16))
	fmt.Println(generateMatrix(25))
}
