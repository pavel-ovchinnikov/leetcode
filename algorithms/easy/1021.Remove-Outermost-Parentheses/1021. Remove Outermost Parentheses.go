package main

import "fmt"

func removeOuterParentheses(s string) string {
	stack := 0
	res := make([]rune, 0, len(s))

	for _, c := range s {
		if c == '(' {
			stack++
		}
		if stack > 1 {
			res = append(res, c)
		}
		if c == ')' {
			stack--
		}
	}
	return string(res)
}

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))
}
