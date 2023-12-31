package main

import (
	"fmt"
)

func insert(intervals [][]int, newInterval []int) [][]int {

	res := make([][]int, 0)

	i := 0

	for ; i < len(intervals) && intervals[i][1] < newInterval[0]; i++ {
		res = append(res, intervals[i])
	}

	for ; i < len(intervals) && intervals[i][0] <= newInterval[1]; i++ {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
	}

	res = append(res, newInterval)

	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}

	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(insert([][]int{
		{1, 2},
		{3, 5},
		{6, 7},
		{8, 10},
		{12, 16},
	}, []int{4, 8}))
	// [1,2],[3,10],[12,16]
}
