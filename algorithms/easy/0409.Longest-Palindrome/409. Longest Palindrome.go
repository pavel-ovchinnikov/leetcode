package main

import (
	"fmt"
)

func longestPalindrome(s string) int {
	letters := make([]int, 52)
	idx := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' {
			idx = int(26 + s[i] - 'a')
		} else {
			idx = int(s[i] - 'A')
		}

		letters[idx]++
	}

	res := 0
	for _, val := range letters {
		res += val / 2 * 2
		if res%2 == 0 && val%2 == 1 {
			res += 1
		}
	}

	return res
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}
