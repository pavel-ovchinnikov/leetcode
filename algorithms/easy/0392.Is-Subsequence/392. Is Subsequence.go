package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	i1, i2 := 0, 0

	for i1 < len(s) && i2 < len(t) {
		if s[i1] == t[i2] {
			i1++
			i2++
			continue
		}

		i2++
	}

	return i1 == len(s)
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}
