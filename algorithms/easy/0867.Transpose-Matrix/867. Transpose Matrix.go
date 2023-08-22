package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	res := make([][]int, n)

	for i := range res {
		res[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res[i][j] = matrix[j][i]
		}
	}

	return res
}

func main() {
	fmt.Println(transpose([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
}
