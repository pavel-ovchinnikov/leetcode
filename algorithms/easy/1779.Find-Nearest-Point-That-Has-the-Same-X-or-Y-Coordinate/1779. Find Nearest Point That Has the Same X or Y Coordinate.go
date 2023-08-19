package main

import (
	"fmt"
	"math"
)

func nearestValidPoint(x int, y int, points [][]int) int {
	index := -1
	dist := math.MaxInt
	for i, point := range points {
		if x != point[0] && y != point[1] {
			continue
		}

		if d := distance(x, y, point[0], point[1]); d < dist {
			index = i
			dist = d
		}
	}
	return index
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func distance(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func main() {
	fmt.Println(nearestValidPoint(3, 4,
		[][]int{
			{1, 2},
			{3, 1},
			{2, 4},
			{2, 3},
			{4, 4}}))
}
