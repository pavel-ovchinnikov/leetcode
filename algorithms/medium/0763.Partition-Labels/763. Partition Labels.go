package main

import (
	"fmt"
)

func partitionLabels(s string) []int {
	hash := map[rune]int{}
	for i, c := range s {
		hash[c] = i
	}

	res := []int{}
	l, r := -1, 0
	for i, c := range s {
		r = max(r, hash[c])
		if r == i {
			res = append(res, r-l)
			l = r
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}
