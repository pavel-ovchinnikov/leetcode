package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func countGoodRectangles(rectangles [][]int) int {
	max, count := 0, 0
	for _, rect := range rectangles {
		maxRect := min(rect[0], rect[1])
		if max < maxRect {
			max = maxRect
			count = 1
		} else if max == maxRect {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countGoodRectangles([][]int{
		{5, 8},
		{3, 9},
		{5, 12},
		{16, 5},
	}))
}
