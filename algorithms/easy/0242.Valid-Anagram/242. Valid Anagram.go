package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	n := len(s)
	arr := make([]int, 26)

	for i := 0; i < n; i++ {
		arr[s[i]-'a']++
		arr[t[i]-'a']--
	}

	for _, v := range arr {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
