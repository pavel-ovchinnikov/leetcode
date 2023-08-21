package main

import (
	"fmt"
	"sort"
)

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	res := make([][]int, 0, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			res = append(res, []int{i, j})
		}
	}

	sort.Slice(res, func(i, j int) bool {
		dist1 := abs(res[i][0]-rCenter) + abs(res[i][1]-cCenter)
		dist2 := abs(res[j][0]-rCenter) + abs(res[j][1]-cCenter)
		return dist1 < dist2
	})

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(allCellsDistOrder(10, 10, 3, 4))
}
