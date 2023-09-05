package main

import "fmt"

func minSteps(s string, t string) int {
	n := len(s)
	count := 0
	letters := make([]int, 26)

	for i := 0; i < n; i++ {
		letters[s[i]-'a'] -= 1
		letters[t[i]-'a'] += 1
	}

	for _, num := range letters {
		if num > 0 {
			count += num
		}
	}

	return count
}

func main() {
	fmt.Println(minSteps("leetcode", "practice"))
}
