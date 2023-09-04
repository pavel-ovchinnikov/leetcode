package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	n := len(s)
	maxPal := ""

	for i := 0; i < n; i++ {
		l, r := extendPalindrome(s, i)
		maxPal = maxPalindrome(maxPal, s[l:r+1])
	}
	return maxPal
}

func extendPalindrome(s string, i int) (l, r int) {
	n := len(s)
	l, r = i, i

	for 0 <= l-1 && s[l-1] == s[i] {
		l--
	}

	for r+1 < n && s[r+1] == s[i] {
		r++
	}

	for 0 <= l-1 && r+1 < n && s[l-1] == s[r+1] {
		l--
		r++
	}

	return
}

func maxPalindrome(s1, s2 string) string {
	if len(s1) < len(s2) {
		return s2
	}
	return s1
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("aa"))
	fmt.Println(longestPalindrome("av"))
}
