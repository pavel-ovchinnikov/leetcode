package main

import "fmt"

func maxVowels(s string, k int) int {
	maxVowel := 0

	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			maxVowel++
		}
	}
	curVowel := maxVowel
	for i := 1; i < len(s)-k+1; i++ {
		// fmt.Println(s[i:i+k], maxVowel, curVowel, string(s[i+k-1]))
		if isVowel(s[i-1]) {
			curVowel--
		}
		if isVowel(s[i+k-1]) {
			curVowel++
		}
		maxVowel = max(maxVowel, curVowel)
	}

	return maxVowel
}

func isVowel(r byte) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	// fmt.Println(maxVowels("abciiidef", 3))
	fmt.Println(maxVowels("weallloveyou", 7))
}
