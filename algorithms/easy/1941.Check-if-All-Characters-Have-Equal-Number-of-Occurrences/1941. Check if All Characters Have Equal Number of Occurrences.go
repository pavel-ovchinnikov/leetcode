package main

import "fmt"

func areOccurrencesEqual(s string) bool {
	letters := make([]int, 26)

	for _, c := range s {
		letters[c-'a']++
	}
	occurrences := 0
	for _, val := range letters {
		switch {
		case val == 0:
			continue
		case occurrences == 0:
			occurrences = val
		case occurrences != val:
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(areOccurrencesEqual("abacbc"))
}
