package main

import "fmt"

// func maxArea(height []int) int {
// 	maxSquare := 0
// 	n := len(height)
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			wight := j - i
// 			square := min(height[i], height[j]) * wight
// 			maxSquare = max(maxSquare, square)
// 		}
// 	}
// 	return maxSquare
// }

func maxArea(height []int) int {
	maxSquare := 0
	n := len(height)
	l, r := 0, n-1
	for l < r {
		wight := r - l
		square := min(height[l], height[r]) * wight
		maxSquare = max(maxSquare, square)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxSquare
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
