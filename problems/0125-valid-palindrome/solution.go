package tasks0125

import "strings"

func isPalindrome(s string) bool {
	runes := []rune(s)

	l, r := 0, len(runes)-1
	for l < r {
		for l < r && !isAlphanumeric(runes[l]) {
			l++
		}
		for l < r && !isAlphanumeric(runes[r]) {
			r--
		}

		if !strings.EqualFold(string(runes[l]), string(runes[r])) {
			return false
		}

		l++
		r--
	}

	return true
}

func isAlphanumeric(c rune) bool {
	return (c >= 'A' && c <= 'Z') ||
		(c >= 'a' && c <= 'z') ||
		(c >= '0' && c <= '9')
}
