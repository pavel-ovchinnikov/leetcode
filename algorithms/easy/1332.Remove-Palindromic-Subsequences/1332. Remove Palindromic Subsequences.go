package main

import "fmt"

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}

	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return 2
		}
		l++
		r--
	}
	return 1
}

func main() {
	fmt.Println(removePalindromeSub("aaaabaaaa"))
}
