package main

import "fmt"

func maxCount(m int, n int, ops [][]int) int {
	row, col := m, n

	for i := 0; i < len(ops); i++ {
		row = Min(row, ops[i][0])
		col = Min(col, ops[i][1])
	}

	return row * col
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxCount(3, 3, [][]int{
		{2, 2},
		{3, 3},
		{3, 3},
		{3, 3},
		{2, 2},
		{3, 3},
		{3, 3},
		{3, 3},
		{2, 2},
		{3, 3},
		{3, 3},
		{3, 3},
	}))
}
