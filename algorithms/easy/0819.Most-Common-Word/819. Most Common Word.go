package main

import (
	"fmt"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	words := split(paragraph)
	blacklist := make(map[string]struct{})
	for _, word := range banned {
		blacklist[word] = struct{}{}
	}

	res := ""
	freq := 0
	for word, count := range words {
		if _, ok := blacklist[word]; ok {
			continue
		}
		if count > freq {
			res = word
			freq = count
		}
	}
	return res
}

func split(s string) map[string]int {
	n := len(s)
	res := make(map[string]int)

	for l := 0; l < n; l++ {
		if ignore(s[l]) {
			continue
		}
		r := l
		for r+1 < n && !ignore(s[r+1]) {
			r++
		}
		res[strings.ToLower(s[l:r+1])]++
		l = r
	}

	return res
}

func ignore(c byte) bool {
	switch c {
	case ' ', '!', '?', '\'', ';', '.', ',':
		return true
	}
	return false
}

func main() {
	fmt.Println(mostCommonWord(
		"Bob hit a ball, the hit BALL flew far after it was hit.",
		[]string{"hit"},
	))
}
