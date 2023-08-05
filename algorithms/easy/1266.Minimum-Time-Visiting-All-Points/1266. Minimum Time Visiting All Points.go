package main

import "fmt"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func seconds(point1, point2 []int) int {
	dx := abs(point1[0] - point2[0])
	dy := abs(point1[1] - point2[1])
	return max(dx, dy)
}

func minTimeToVisitAllPoints(points [][]int) int {
	res := 0
	for i := 1; i < len(points); i++ {
		fmt.Println(seconds(points[i-1], points[i]))
		res += seconds(points[i-1], points[i])
	}
	return res
}

func main() {
	fmt.Println(minTimeToVisitAllPoints([][]int{
		{1, 1},
		{3, 4},
		{-1, 0},
	}))
}
