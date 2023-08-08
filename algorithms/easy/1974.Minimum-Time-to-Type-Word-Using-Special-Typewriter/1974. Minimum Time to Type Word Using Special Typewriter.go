package main

import "fmt"

func minTimeToType(word string) int {
	sec := 0
	cur := 'a'
	for _, c := range word {
		if c-cur != 0 {
			diff := minTurn(abs(int(c - cur)))
			sec += diff
		}
		sec++
		cur = c
	}
	return sec
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minTurn(a int) int {
	if a > 13 {
		return 26 - a
	}
	return a
}

func main() {
	fmt.Println(minTimeToType("abc"))
	fmt.Println(minTimeToType("bza"))
	fmt.Println(minTimeToType("zjpc"))
}
