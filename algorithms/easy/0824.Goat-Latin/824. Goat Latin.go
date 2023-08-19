package main

import (
	"fmt"
	"strings"
)

func toGoatLatin(sentence string) string {
	res := ""
	words := strings.Split(sentence, " ")

	for i, word := range words {
		goatWord := ""

		switch word[0] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			goatWord = word + "ma"
		default:
			goatWord = word[1:] + string(word[0]) + "ma"
		}

		res = res + goatWord + strings.Repeat("a", i+1) + " "
	}

	return res[:len(res)-1]
}

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))
}
