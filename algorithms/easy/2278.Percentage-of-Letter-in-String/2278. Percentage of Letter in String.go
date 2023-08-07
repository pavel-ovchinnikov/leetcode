package main

import (
	"fmt"
)

func percentageLetter(s string, letter byte) int {
	count := 0

	for _, c := range s {
		if c == rune(letter) {
			count++
		}
	}
	return count * 100 / len(s)
}

// func percentageLetter(s string, letter byte) int {
// 	return strings.Count(s, string(letter)) * 100 / len(s)
// }

func main() {
	fmt.Println(percentageLetter("foobar", 'o'))
}
