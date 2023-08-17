package main

import (
	"fmt"
	"math"
)

func countPoints(points [][]int, queries [][]int) []int {
	result := make([]int, len(queries))
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(queries); j++ {
			if distance(points[i][0], points[i][1], queries[j][0], queries[j][1]) <= float64(queries[j][2]) {
				result[j] += 1
			}
		}
	}
	return result
}

func distance(x, y, px, py int) float64 {
	return math.Sqrt(math.Pow(float64(x)-float64(px), 2) + math.Pow(float64(y)-float64(py), 2))
}

func main() {
	fmt.Println(countPoints(
		[][]int{
			{1, 3},
			{3, 3},
			{5, 3},
			{2, 2}},
		[][]int{
			{2, 3, 1},
			{4, 3, 1},
			{1, 1, 2},
		}))
}
