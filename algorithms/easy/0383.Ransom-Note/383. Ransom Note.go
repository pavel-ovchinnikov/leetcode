package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	n := 26
	letters := make([]int, n)

	for _, c := range magazine {
		letters[c-'a']++
	}

	for _, c := range ransomNote {
		letters[c-'a']--
	}

	for _, v := range letters {
		if v < 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canConstruct("aa", "aadf"))
}
