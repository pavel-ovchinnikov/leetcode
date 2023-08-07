package main

import "fmt"

func minimizedStringLength(s string) int {
	letters := make([]int, 26)
	res := 0
	for _, c := range s {
		letters[c-'a']++
		if letters[c-'a'] == 1 {
			res++
		}
	}

	return res
}

func main() {
	fmt.Println(minimizedStringLength("aaab"))
}
