package main

import "fmt"

func winnerOfGame(colors string) bool {
	countA, countB := 0, 0
	l, r := 0, 0
	n := len(colors)
	for l < n {
		r = l
		for r+1 < n && colors[l] == colors[r+1] {
			r++
		}

		d := r - l + 1

		if d >= 3 {
			if colors[l] == 'A' {
				countA += d - 2
			} else {
				countB += d - 2
			}
		}

		l = r + 1
	}
	// fmt.Println(countA, countB)
	return countA > countB
}

func main() {
	fmt.Println(winnerOfGame("AAABABB"))
	fmt.Println(winnerOfGame("BBAAABBABBABB"))
	// BBAAABBABBABB
}
