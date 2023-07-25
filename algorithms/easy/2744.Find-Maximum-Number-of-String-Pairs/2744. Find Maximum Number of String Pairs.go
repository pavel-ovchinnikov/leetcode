package main

import (
	"fmt"
	"sort"
)

func maximumNumberOfStringPairs(words []string) int {
	count, res := make(map[string]int), 0
	for _, w := range words {
		bytes := []byte(w)

		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})

		count[string(bytes)]++
		if count[string(bytes)] >= 2 {
			res++
		}
	}

	return res
}

func main() {
	fmt.Println(maximumNumberOfStringPairs([]string{"cd", "ac", "dc", "ca", "zz"}))
}
