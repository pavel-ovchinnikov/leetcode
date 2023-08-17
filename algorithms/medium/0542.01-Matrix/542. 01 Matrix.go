package main

import (
	"fmt"
)

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	D := m + n // max distance from cells

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				continue
			}
			a, b := D, D
			if i-1 >= 0 {
				a = mat[i-1][j] + 1
			}
			if j-1 >= 0 {
				b = mat[i][j-1] + 1
			}
			mat[i][j] = min(a, b)
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if mat[i][j] == 0 {
				continue
			}
			a, b := D, D
			if i+1 < m {
				a = mat[i+1][j] + 1
			}
			if j+1 < n {
				b = mat[i][j+1] + 1
			}
			mat[i][j] = min(mat[i][j], min(a, b))
		}
	}
	return mat
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(updateMatrix([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1}}))
}
