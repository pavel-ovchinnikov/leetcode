package main

import "fmt"

func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}
	res := 0
	isWord := false
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			if !isWord {
				isWord = true
				res++
			}
		} else {
			isWord = false
		}
	}
	return res
}

func main() {
	fmt.Println(countSegments("Hello, my name is John"))
}
