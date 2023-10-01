package main

import "fmt"

func reverseWords(s string) string {
	res := []byte(s)

	l, r := 0, 0
	for l < len(s) {
		if s[l] == ' ' {
			l++
			continue
		}

		for r = l + 1; r < len(s) && s[r] != ' '; r++ {
		}

		r = Max(l, r)
		reverse(res[l:r])

		l = r + 1
	}

	return string(res)
}

func reverse(arr []byte) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	// fmt.Println(reverseWords("  Let's  take LeetCode contest "))

	fmt.Println(reverseWords("vector<string> split (string s, char delimiter)"))
}
