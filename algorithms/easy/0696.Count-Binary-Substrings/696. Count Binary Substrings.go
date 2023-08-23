package main

import "fmt"

func countBinarySubstrings(s string) int {
	res, cur, prev := 0, 1, 0

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cur++
		} else {
			res += min(prev, cur)
			prev, cur = cur, 1
		}
	}

	return res + min(prev, cur)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(countBinarySubstrings("00110011"))
}
