package main

import (
	"fmt"
	"sort"
)

func reorganizeString(s string) string {
	n := len(s)
	freq := make(map[rune]int)
	for _, c := range s {
		freq[c]++
	}

	letters := make([]rune, 0, len(freq))
	for c := range freq {
		letters = append(letters, c)
	}

	sort.Slice(letters, func(i, j int) bool {
		return freq[letters[j]] < freq[letters[i]]
	})

	if freq[letters[0]] > (n+1)/2 {
		return ""
	}

	res := make([]rune, len(s))
	i := 0
	for _, ch := range letters {
		for j := 0; j < freq[ch]; j++ {
			if i >= len(s) {
				i = 1
			}
			res[i] = ch
			i += 2
		}
	}

	return string(res)
}

func main() {
	fmt.Println(reorganizeString("a"))
	fmt.Println(reorganizeString("aa"))
	fmt.Println(reorganizeString("aab"))
	fmt.Println(reorganizeString("aabb"))
	fmt.Println(reorganizeString("caaabb"))
	fmt.Println(reorganizeString("aaab"))

	fmt.Println(reorganizeString("vlovv"))
}
