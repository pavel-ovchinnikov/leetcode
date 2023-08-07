package main

import "fmt"

func prefixCount(words []string, pref string) int {
	res := 0
	n := len(pref)
	for _, w := range words {
		if len(w) >= n && w[:n] == pref {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(prefixCount([]string{"pay", "attention", "practice", "attend"}, "at"))
}
