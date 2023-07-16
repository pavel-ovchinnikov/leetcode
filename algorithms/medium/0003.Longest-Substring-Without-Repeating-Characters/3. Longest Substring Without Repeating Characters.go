package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	res := 0
	size := 0
	last := make(map[rune]int)

	for _, c := range s {
		last[c] = -1
	}

	for i, c := range s {
		size = Min(size+1, i-last[c])
		last[c] = i
		res += size - Min(res, size)
	}
	return res
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(" "))
}
