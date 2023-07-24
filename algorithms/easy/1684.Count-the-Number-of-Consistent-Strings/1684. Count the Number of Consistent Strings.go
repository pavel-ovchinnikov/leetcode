package main

import "fmt"

func countConsistentStrings(allowed string, words []string) int {
	mm := make(map[rune]bool)
	res := 0

	for _, c := range allowed {
		mm[c] = true
	}

	for _, w := range words {
		f := true
		for _, c := range w {
			if _, contains := mm[c]; !contains {
				f = false
				break
			}
		}
		if f {
			res += 1
		}
	}

	return res
}

func main() {
	fmt.Println(countConsistentStrings("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))
	fmt.Println(countConsistentStrings("abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}))
	fmt.Println(countConsistentStrings("cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}))
}
