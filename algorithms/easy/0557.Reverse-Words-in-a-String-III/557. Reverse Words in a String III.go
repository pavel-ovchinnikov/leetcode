package main

import "fmt"

func reverseWords(s string) string {
	result := []byte(s)
	n := len(result)
	start, end := 0, 0
	lastSpaceIndex := -1

	for i := 0; i <= n; i++ {
		if i == n || result[i] == ' ' {
			start, end = lastSpaceIndex+1, i-1
			for i, j := start, end; i < j; i, j = i+1, j-1 {
				result[i], result[j] = result[j], result[i]
			}
			lastSpaceIndex = i
		}
	}

	return string(result)
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}
