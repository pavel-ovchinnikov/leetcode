package main

import "fmt"

func lengthOfLastWord(s string) int {
	count := 0
	i := len(s) - 1

	for ; i >= 0 && s[i] == ' '; i -= 1 {
	}
	for ; i >= 0 && s[i] != ' '; i -= 1 {
		count += 1
	}

	return count
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}
