package main

import (
	"fmt"
	"unicode"
)

func reverseVowels(s string) string {
	res := []rune(s)
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !isVowel(res[l]) {
			l++
		}

		for l < r && !isVowel(res[r]) {
			r--
		}
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return string(res)
}

func isVowel(c rune) bool {
	c = unicode.ToLower(c)
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func main() {
	fmt.Println(reverseVowels("hello"))
}
