package main

import "fmt"

func countGoodSubstrings(s string) int {
	res := 0
	for i := 2; i < len(s); i++ {
		if !(s[i] == s[i-1] || s[i-1] == s[i-2] || s[i-2] == s[i]) {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(countGoodSubstrings("xyzzaz"))
}
