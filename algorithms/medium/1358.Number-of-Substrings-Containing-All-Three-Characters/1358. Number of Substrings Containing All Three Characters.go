package main

import (
	"fmt"
	"math"
)

func numberOfSubstrings(s string) int {
	n := len(s)
	res := 0
	indexA, indexB, indexC := -1, -1, -1
	for i, c := range s {
		if c == 'a' && indexA == -1 {
			indexA = i
			continue
		}

		if c == 'b' && indexB == -1 {
			indexB = i
			continue
		}

		if c == 'c' && indexC == -1 {
			indexC = i
			continue
		}
	}

	if min(indexA, indexB, indexC) == -1 {
		return 0
	}

	nextIndex := func(startIndex int, c byte) int {
		res := n
		for i := startIndex; i < n; i++ {
			if s[i] == c {
				return i
			}
		}
		return res
	}

	for index := max(indexA, indexB, indexC); index < n; index = max(indexA, indexB, indexC) {
		res += n - index
		minIndex := min(indexA, indexB, indexC)
		switch s[minIndex] {
		case 'a':
			indexA = nextIndex(indexA+1, 'a')
		case 'b':
			indexB = nextIndex(indexB+1, 'b')
		case 'c':
			indexC = nextIndex(indexC+1, 'c')
		}
	}

	return res
}

func max(nums ...int) int {
	maxVal := math.MinInt
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}
	return maxVal
}

func min(nums ...int) int {
	minVal := math.MaxInt
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}

func main() {
	fmt.Println(numberOfSubstrings("abcabc"))
	fmt.Println(numberOfSubstrings("aaacb"))
	fmt.Println(numberOfSubstrings("abc"))
	fmt.Println(numberOfSubstrings("abab"))
}
