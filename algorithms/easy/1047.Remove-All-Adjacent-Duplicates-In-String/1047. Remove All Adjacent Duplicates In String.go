package main

import "fmt"

func removeDuplicates(s string) string {
	stack := make([]rune, 0)

	for _, c := range s {
		if len(stack) > 0 && c == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}
