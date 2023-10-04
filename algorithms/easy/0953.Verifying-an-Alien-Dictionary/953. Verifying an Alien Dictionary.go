package main

import (
	"fmt"
	"sort"
)

func isAlienSorted(words []string, order string) bool {
	m := make(map[rune]rune)
	for i, b := range order {
		m[b] = rune('a' + byte(i))
	}
	for i, s := range words {
		runes := []rune(s)
		for i, r := range runes {
			runes[i] = m[r]
		}
		words[i] = string(runes)
	}
	return sort.StringsAreSorted(words)
}

func main() {
	fmt.Println(isAlienSorted(
		[]string{"hello", "leetcode"},
		"hlabcdefgijkmnopqrstuvwxyz",
	))
}
