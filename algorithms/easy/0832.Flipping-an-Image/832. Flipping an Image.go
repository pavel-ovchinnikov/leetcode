package main

import "fmt"

func flipAndInvertImage(image [][]int) [][]int {
	n := len(image)
	m := n / 2
	if n%2 != 0 {
		m = n/2 + 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			tmp := image[i][j]
			image[i][j] = (image[i][n-1-j] + 1) % 2
			image[i][n-1-j] = (tmp + 1) % 2
		}
	}
	return image
}
func main() {
	fmt.Println(flipAndInvertImage([][]int{
		{1, 1, 0},
		{1, 0, 1},
		{0, 0, 0},
	}))
}
