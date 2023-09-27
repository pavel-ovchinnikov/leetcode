package main

import "fmt"

func matrixReshape(mat [][]int, r int, c int) [][]int {
	n, m := len(mat), len(mat[0])

	if n*m != r*c {
		return mat
	}

	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			count := i*m + j
			row := count / c
			col := count % c
			res[row][col] = mat[i][j]
		}
	}

	return res
}

func main() {
	fmt.Println(matrixReshape([][]int{
		{1, 2},
		{3, 4},
	}, 1, 4))
}
