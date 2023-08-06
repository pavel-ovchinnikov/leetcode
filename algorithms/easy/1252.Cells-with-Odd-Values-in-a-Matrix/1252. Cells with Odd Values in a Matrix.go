package main

import "fmt"

// func oddCells(m int, n int, indices [][]int) int {
// 	grid := make([][]int, m)
// 	for i := 0; i < m; i++ {
// 		grid[i] = make([]int, n)
// 	}

// 	for _, index := range indices {
// 		for k := 0; k < m; k++ {
// 			grid[k][index[1]]++
// 		}
// 		for k := 0; k < n; k++ {
// 			grid[index[0]][k]++
// 		}
// 	}
// 	res := 0
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if grid[i][j]%2 == 1 {
// 				res++
// 			}
// 		}
// 	}

// 	return res
// }

func oddCells(m int, n int, indices [][]int) int {
	res := 0
	rows, cols := make([]int, m), make([]int, n)
	for _, idx := range indices {
		rows[idx[0]]++
		cols[idx[1]]++
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res += (rows[i] + cols[j]) & 1
		}
	}
	return res
}

func main() {
	fmt.Println(oddCells(3, 2, [][]int{{0, 1}, {1, 1}}))
}
