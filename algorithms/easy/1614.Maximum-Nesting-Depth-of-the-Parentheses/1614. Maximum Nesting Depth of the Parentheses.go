package main

import "fmt"

func maxDepth(s string) int {
	stack := 0
	res := 0

	for i := range s {
		switch s[i] {
		case '(':
			stack++
			if stack > res {
				res = stack
			}
		case ')':
			stack--
		}
	}
	return res
}

func main() {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1"))
}
