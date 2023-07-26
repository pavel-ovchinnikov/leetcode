package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mergeAlternately(word1 string, word2 string) string {
	n1, n2 := len(word1), len(word2)
	res := make([]byte, 0, n1+n2)
	n := min(n1, n2)

	for i := 0; i < n; i++ {
		res = append(res, word1[i])
		res = append(res, word2[i])
	}

	if n2 > n1 {
		for i := n; i < n2; i++ {
			res = append(res, word2[i])
		}
	} else {
		for i := n; i < n1; i++ {
			res = append(res, word1[i])
		}
	}

	return string(res[:n1+n2])
}

func main() {
	fmt.Println(mergeAlternately("abc", "pqr"))
	fmt.Println(mergeAlternately("ab", "pqrs"))
	fmt.Println(mergeAlternately("abcd", "pq"))
}
