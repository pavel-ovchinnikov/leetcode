package main

import "fmt"

func largeGroupPositions(s string) [][]int {
	res := [][]int{}
	n := len(s)

	for l := 0; l < n; l++ {
		r := l
		for r+1 < n && s[l] == s[r+1] {
			r++
		}
		if r-l >= 2 {
			res = append(res, []int{l, r})
		}
		l = r
	}
	return res
}

func main() {
	fmt.Println(largeGroupPositions("abbxxxxzzy"))
	fmt.Println(largeGroupPositions("abs"))
	fmt.Println(largeGroupPositions("abcdddeeeeaabbbcd"))
}
