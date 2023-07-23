package main

import "fmt"

func truncateSentence(s string, k int) string {
	count := 0
	for i, c := range s {
		if c == ' ' {
			count++
		}

		if count == k {
			return s[:i]
		}

	}
	return s
}

func main() {
	fmt.Println(truncateSentence("Hello how are you Contestant", 4))
}
