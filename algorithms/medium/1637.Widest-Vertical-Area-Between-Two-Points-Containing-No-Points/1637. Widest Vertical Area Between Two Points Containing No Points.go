package main

import (
	"fmt"
	"sort"
)

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	maxWidth := 0
	for i := 1; i < len(points); i++ {
		if points[i][0] != points[i-1][0] {
			maxWidth = max(maxWidth, points[i][0]-points[i-1][0])
		}
	}

	return maxWidth
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(maxWidthOfVerticalArea([][]int{
		{3, 1},
		{9, 0},
		{1, 0},
		{1, 4},
		{5, 3},
		{8, 8},
	}))
}
