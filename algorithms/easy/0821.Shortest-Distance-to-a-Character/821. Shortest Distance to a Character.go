package main

import (
	"fmt"
	"math"
)

func shortestToChar(s string, c byte) []int {
	n := len(s)

	res := make([]int, n)
	index := -1
	for i := 0; i < n; i++ {
		if s[i] == c {
			index = i
			continue
		}

		if index == -1 {
			res[i] = math.MaxInt
			continue
		}

		res[i] = i - index
	}

	index = -1
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			index = i
			continue
		}

		if index == -1 {
			continue
		}

		res[i] = min(res[i], index-i)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(shortestToChar("loveleetcode", 'e'))
}
