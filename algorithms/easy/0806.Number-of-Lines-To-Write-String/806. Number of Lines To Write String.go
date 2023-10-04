package main

import "fmt"

func numberOfLines(widths []int, s string) []int {
	row, col := 0, 0

	for _, c := range s {
		w := widths[c-'a']
		if col+w > 100 {
			row++
			col = 0
		}
		col += w
	}

	if col > 0 {
		row++
	}

	return []int{row, col}
}

func main() {
	fmt.Println(numberOfLines(
		[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		"abcdefghijklmnopqrstuvwxyz",
	))
}
