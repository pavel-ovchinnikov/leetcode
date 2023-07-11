package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		n := len(stack) - 1
		prev := stack[n]
		stack = stack[:n]

		if !((c == ')' && prev == '(') ||
			(c == ']' && prev == '[') ||
			(c == '}' && prev == '{')) {
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
}
