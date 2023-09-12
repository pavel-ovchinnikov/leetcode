package main

import (
	"fmt"
	"sort"
)

func minDeletions(s string) int {
	freq := make([]int, 26)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}
	sort.Slice(freq, func(i, j int) bool {
		return freq[i] > freq[j]
	})

	count := 0
	prev := freq[0]
	for i := 1; i < len(freq); i++ {
		if freq[i] == 0 {
			break
		}

		if prev <= 1 {
			count += freq[i]
			prev = 0
			continue
		}

		if prev <= freq[i] {
			count += freq[i] - (prev - 1)
			prev--
			continue
		}

		prev = freq[i]
	}

	return count
}

func main() {
	fmt.Println(minDeletions("aabbbi"))
	fmt.Println(minDeletions("aaabbbcc"))
	fmt.Println(minDeletions("ceabaacb"))
	fmt.Println(minDeletions("abcd"))
}
