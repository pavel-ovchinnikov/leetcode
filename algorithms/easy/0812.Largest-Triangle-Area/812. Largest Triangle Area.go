package main

import (
	"fmt"
	"math"
)

func largestTriangleArea(points [][]int) float64 {
	res := 0.0
	n := len(points)
	for i := 0; i < n-2; i++ {
		for j := 0; j < n-1; j++ {
			for k := 0; k < n; k++ {
				res = math.Max(res, area(points[i], points[j], points[k]))
			}
		}
	}
	return res
}

func area(p1, p2, p3 []int) float64 {
	return math.Abs(float64(
		p1[0]*(p2[1]-p3[1])+
			p2[0]*(p3[1]-p1[1])+
			p3[0]*(p1[1]-p2[1])) / 2.0)
}

func main() {
	fmt.Println(largestTriangleArea([][]int{
		{0, 0},
		{0, 1},
		{1, 0},
		{0, 2},
		{2, 0},
	}))
}
