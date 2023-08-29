package main

import (
	"fmt"
	"sort"
)

func diagonalSort(mat [][]int) [][]int {
	n, m := len(mat), len(mat[0])
	hash := make(map[int][]int)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			hash[i-j] = append(hash[i-j], mat[i][j])
		}
	}

	for _, nums := range hash {
		sort.Ints(nums)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			mat[i][j] = hash[i-j][0]
			hash[i-j] = hash[i-j][1:]
		}
	}

	return mat
}

func main() {
	fmt.Println(diagonalSort([][]int{
		{3, 3, 1, 1},
		{2, 2, 1, 2},
		{1, 1, 1, 2}}))
}
