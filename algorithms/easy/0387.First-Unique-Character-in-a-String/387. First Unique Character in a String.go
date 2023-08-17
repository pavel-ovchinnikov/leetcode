package main

import "fmt"

func firstUniqChar(s string) int {
	if len(s) == 1 {
		return 0
	}

	letters := make([]int, 26)
	for i := range letters {
		letters[i] = -1
	}

	for i, c := range s {
		if letters[c-'a'] == -2 {
			continue
		}

		if letters[c-'a'] == -1 {
			letters[c-'a'] = i
			continue
		}

		if letters[c-'a'] >= 0 {
			letters[c-'a'] = -2
			continue
		}
	}

	minIndex := len(s)
	for i := 0; i < len(letters); i++ {
		if letters[i] == -1 || letters[i] == -2 {
			continue
		}

		minIndex = min(minIndex, letters[i])
	}

	if minIndex == len(s) {
		return -1
	}

	return minIndex
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(firstUniqChar("dddccdbba"))
}
