package main

import "fmt"

func canBeTypedWords(text string, brokenLetters string) int {
	letters := make(map[byte]bool)
	count := 1

	for _, v := range brokenLetters {
		letters[byte(v)] = true
	}

	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			count++
			continue
		}

		if letters[text[i]] {
			count--
			for i+1 < len(text) && text[i+1] != ' ' {
				i++
			}
		}

	}

	return count
}

func main() {
	fmt.Println(canBeTypedWords("Hello word w", "ad"))
}
