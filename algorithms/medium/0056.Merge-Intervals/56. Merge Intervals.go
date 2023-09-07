package main

import (
	"fmt"
	"sort"
)

// func merge(intervals [][]int) [][]int {
// 	sort.Slice(intervals, func(i, j int) bool {
// 		return intervals[i][0] < intervals[j][0]
// 	})

// 	i, j := 0, 1

// 	for j < len(intervals) {
// 		interval := intervals[j]
// 		if intervals[i][1] >= interval[0] {
// 			intervals[i][1] = max(intervals[i][1], interval[1])
// 			intervals = append(intervals[:j], intervals[j+1:]...)
// 		} else {
// 			i = j
// 			j++
// 		}
// 	}
// 	return intervals
// }

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{intervals[0]}
	for j := 1; j < len(intervals); j++ {
		interval := intervals[j]
		if res[len(res)-1][1] >= interval[0] {
			res[len(res)-1][1] = max(res[len(res)-1][1], interval[1])
		} else {
			res = append(res, interval)
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
	fmt.Println(merge([][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}))
	fmt.Println(merge([][]int{
		{1, 4},
		{4, 5},
	}))
	fmt.Println(merge([][]int{
		{1, 4},
		{2, 3},
	}))
}
