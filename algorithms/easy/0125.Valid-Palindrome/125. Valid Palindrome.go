package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for ; l < r; l, r = l+1, r-1 {
		for l < r && !(isLetterOrIsDigit(rune(s[l]))) {
			l++
		}

		for l < r && !(isLetterOrIsDigit(rune(s[r]))) {
			r--
		}

		if l == r {
			break
		}

		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
	}
	return true
}

func isLetterOrIsDigit(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}
