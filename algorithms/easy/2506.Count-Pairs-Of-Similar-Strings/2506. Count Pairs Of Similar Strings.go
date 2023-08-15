package main

import "fmt"

func similarPairs(words []string) int {
	res := 0
	hash := map[[26]bool]int{}
	for _, word := range words {
		letters := [26]bool{}
		for _, c := range word {
			letters[c-'a'] = true
		}
		hash[letters]++
	}
	for _, count := range hash {
		res += count * (count - 1) / 2
	}
	return res
}

func main() {
	fmt.Println(similarPairs([]string{"aba", "aabb", "abcd", "bac", "aabc"}))
}
