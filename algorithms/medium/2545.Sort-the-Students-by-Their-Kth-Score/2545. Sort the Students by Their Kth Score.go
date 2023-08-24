package main

import (
	"fmt"
	"sort"
)

func sortTheStudents(score [][]int, k int) [][]int {
	sort.SliceStable(score, func(i, j int) bool {
		return score[i][k] > score[j][k]
	})

	return score
}

func main() {
	grid := [][]int{
		{10, 6, 9, 1},
		{7, 5, 11, 2},
		{4, 8, 3, 15},
	}

	fmt.Println(sortTheStudents(grid, 2))
}
