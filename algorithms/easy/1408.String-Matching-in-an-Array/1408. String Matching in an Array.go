package main

import (
	"fmt"
	"sort"
	"strings"
)

func stringMatching(words []string) []string {
	res := []string{}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if strings.Contains(words[j], words[i]) {
				res = append(res, words[i])
				break
			}
		}
	}

	return res
}

func main() {
	fmt.Println(stringMatching([]string{"mass", "as", "hero", "superhero"}))
}
