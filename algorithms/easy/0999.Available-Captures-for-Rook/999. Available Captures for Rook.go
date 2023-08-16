package main

import "fmt"

func numRookCaptures(board [][]byte) int {
	res := 0

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'R' {
				for k := i + 1; k < len(board); k++ {
					if board[k][j] == 'p' {
						res++
						break
					}

					if board[k][j] == 'B' {
						break
					}
				}

				for k := i - 1; k >= 0; k-- {
					if board[k][j] == 'p' {
						res++
						break
					}

					if board[k][j] == 'B' {
						break
					}
				}

				for k := j + 1; k < len(board[0]); k++ {
					if board[i][k] == 'p' {
						res++
						break
					}

					if board[i][k] == 'B' {
						break
					}
				}

				for k := j - 1; k >= 0; k-- {
					if board[i][k] == 'p' {
						res++
						break
					}

					if board[i][k] == 'B' {
						break
					}
				}
				return res
			}
		}

	}

	return res
}

func main() {
	fmt.Println(numRookCaptures([][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		{'.', 'p', 'B', 'R', 'B', 'p', '.', '.'},
		{'.', 'p', 'p', 'B', 'p', 'p', '.', '.'},
		{'.', 'p', 'p', 'p', 'p', 'p', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}))
}
