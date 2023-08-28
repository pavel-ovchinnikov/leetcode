package main

import "fmt"

func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}

	for i, word := range words {
		if s[i] != word[0] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAcronym([]string{"alice", "bob", "charlie"}, "abc"))
}
