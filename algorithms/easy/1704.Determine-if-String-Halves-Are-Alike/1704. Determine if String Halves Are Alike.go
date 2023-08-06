package main

import "fmt"

func halvesAreAlike(s string) bool {
	res := 0
	half := len(s) / 2
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			if i < half {
				res++
			} else {
				res--
			}
		}
	}
	return res == 0
}

func main() {
	fmt.Println(halvesAreAlike("book"))
}
