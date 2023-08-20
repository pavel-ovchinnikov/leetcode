package main

import (
	"fmt"
	"unicode"
)

func findWords(words []string) []string {
	res := make([]string, 0, len(words))

	for _, word := range words {
		if isOneRow0fAmericanKeyboard(word) {
			res = append(res, word)
		}
	}

	return res
}

func isOneRow0fAmericanKeyboard(word string) bool {
	flag := -1
	for _, c := range word {

		switch unicode.ToLower(c) {
		case 'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p':
			if flag == -1 {
				flag = 1
			}
			if flag != 1 {
				return false
			}
		case 'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l':
			if flag == -1 {
				flag = 2
			}
			if flag != 2 {
				return false
			}
		case 'z', 'x', 'c', 'v', 'b', 'n', 'm':
			if flag == -1 {
				flag = 3
			}
			if flag != 3 {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}
