package main

import "fmt"

func removeDuplicateLetters(s string) string {
	remaining := map[rune]int{}
	for _, r := range s {
		remaining[r]++
	}
	stackContains := map[rune]bool{}
	stack := []rune{}

	for _, r := range s {
		remaining[r]--
		if stackContains[r] {
			continue
		}
		for len(stack) > 0 && r < stack[len(stack)-1] && remaining[stack[len(stack)-1]] > 0 {
			stackContains[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, r)
		stackContains[r] = true
	}
	return string(stack)
}

func main() {
	fmt.Println(removeDuplicateLetters("cbacdcbc"))
}
