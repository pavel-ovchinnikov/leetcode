package main

import (
	"fmt"
	"sort"
)

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])

	for i := 0; i < n; i++ {
		index := sort.SearchInts(matrix[i], target)
		if index < m && matrix[i][index] == target {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 16))
}
