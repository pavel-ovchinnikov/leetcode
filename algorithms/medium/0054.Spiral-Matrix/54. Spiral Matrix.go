package main

import "fmt"

const (
	Left int = iota
	Right
	Up
	Down
)

func spiralOrder(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	res := make([]int, n*m)
	top, bottom := 0, n-1
	left, right := 0, m-1
	index := 0
	direction := Right

	for index < n*m {
		switch direction {
		case Right:
			for i := left; i <= right; i++ {
				res[index] = matrix[top][i]
				index++
			}
			top++
			direction = Down
		case Down:
			for i := top; i <= bottom; i++ {
				res[index] = matrix[i][right]
				index++
			}
			right--
			direction = Left
		case Left:
			for i := right; i >= left; i-- {
				res[index] = matrix[bottom][i]
				index++
			}
			bottom--
			direction = Up
		case Up:
			for i := bottom; i >= top; i-- {
				res[index] = matrix[i][left]
				index++
			}
			left++
			direction = Right
		}
	}

	return res
}

func main() {
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}))
	fmt.Println([]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7})
}
