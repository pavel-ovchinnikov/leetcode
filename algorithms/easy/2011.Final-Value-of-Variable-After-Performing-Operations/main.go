package main

import "fmt"

func finalValueAfterOperations(operations []string) int {
	result := 0

	for _, val := range operations {
		if val[1] == '+' {
			result++
		} else {
			result--
		}
	}

	return result
}

func main() {
	fmt.Println(finalValueAfterOperations([]string{"--X", "X++", "X++"}))
}
