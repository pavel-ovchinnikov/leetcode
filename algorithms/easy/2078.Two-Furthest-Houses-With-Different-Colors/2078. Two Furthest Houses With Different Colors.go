package main

import "fmt"

func maxDistance(colors []int) int {
	n := len(colors)
	l, r := 0, n-1

	for colors[l] == colors[n-1] {
		l++
	}
	for colors[0] == colors[r] {
		r--
	}

	return max(n-1-l, r)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(maxDistance([]int{1, 1, 1, 3, 2, 1}))
}
