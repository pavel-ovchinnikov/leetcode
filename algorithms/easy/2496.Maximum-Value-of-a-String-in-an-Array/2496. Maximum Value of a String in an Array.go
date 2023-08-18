package main

import (
	"fmt"
	"math"
	"strconv"
)

func maximumValue(strs []string) int {
	res := math.MinInt
	for _, word := range strs {
		for j, c := range word {
			if c <= '9' && j == len(word)-1 {
				a, _ := strconv.Atoi(word)
				res = max(res, a)
				break
			}

			if c >= 'a' {
				res = max(res, len(word))
				break
			}

		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(maximumValue([]string{"alic3", "bob", "3", "4", "00000"}))
	fmt.Println(maximumValue([]string{"1", "01", "001", "0001"}))
}
