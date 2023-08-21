package main

import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	return strings.Contains(s[1:n]+s[0:n-1], s)
}

// abcabc
// bcabcabcab

func main() {
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
}
