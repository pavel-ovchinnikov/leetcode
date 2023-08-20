package main

import "fmt"

func countWords(words1 []string, words2 []string) int {
	res := 0
	hash := make(map[string]int)

	for _, word := range words1 {
		hash[word]++
	}

	for key, val := range hash {
		if val != 1 {
			delete(hash, key)
		}
	}

	for _, word := range words2 {
		hash[word]--
	}

	for _, val := range hash {
		if val == 0 {
			res++
		}
	}

	return res
}

func main() {
	fmt.Println(countWords(
		[]string{"leetcode", "is", "amazing", "as", "is"},
		[]string{"amazing", "leetcode", "is"}))
}
