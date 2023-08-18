package main

import "fmt"

func luckyNumbers(matrix [][]int) []int {
	res := []int{}
	n, m := len(matrix), len(matrix[0])
	cols := make([]int, len(matrix[0]))

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			cols[j] = max(cols[j], matrix[i][j])
		}
	}

	for i := 0; i < n; i++ {
		minIndex := 0
		minVal := matrix[i][minIndex]
		for j := 0; j < m; j++ {
			if matrix[i][j] < minVal {
				minVal = matrix[i][j]
				minIndex = j
			}
		}
		if minVal == cols[minIndex] {
			res = append(res, minVal)
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(luckyNumbers([][]int{
		{3, 7, 8},
		{9, 11, 13},
		{15, 16, 17},
	}))
}
