package main

import "fmt"

func checkIfPangram(sentence string) bool {
	store := make(map[rune]struct{})
	for _, ch := range sentence {
		if _, ok := store[ch]; !ok {
			store[ch] = struct{}{}
		}
	}
	return len(store) == 26
}

func main() {
	fmt.Println(checkIfPangram("leetcode"))
	fmt.Println(checkIfPangram("thequickbrownfoxjumpsoverthelazydog"))
}
