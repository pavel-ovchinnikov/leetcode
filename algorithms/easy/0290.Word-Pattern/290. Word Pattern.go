package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	patternMap := make(map[byte]string, len(pattern))
	sMap := make(map[string]byte, len(s))
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	for i, word := range words {
		_, ok1 := patternMap[pattern[i]]
		_, ok2 := sMap[word]

		if !ok1 && !ok2 {
			patternMap[pattern[i]] = word
			sMap[word] = pattern[i]
		} else if ok2 {
			if sMap[word] != pattern[i] {
				return false
			}
		} else if ok1 {
			if patternMap[pattern[i]] != word {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
}
