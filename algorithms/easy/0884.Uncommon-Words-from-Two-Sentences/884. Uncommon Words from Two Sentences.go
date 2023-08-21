package main

import (
	"fmt"
	"strings"
)

func uncommonFromSentences(s1 string, s2 string) []string {
	res := []string{}
	words := make(map[string]int)

	for _, word := range strings.Split(s1, " ") {
		words[word]++
	}

	for _, word := range strings.Split(s2, " ") {
		words[word]++
	}

	for word, count := range words {
		if count == 1 {
			res = append(res, word)
		}
	}

	return res
}

func main() {
	fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
}
