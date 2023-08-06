package main

import "fmt"

func reversePrefix(word string, ch byte) string {
	n := len(word)
	res := make([]byte, n)

	idx := 0

	for idx < n && word[idx] != ch {
		idx++
	}
	if idx >= n {
		return word
	}

	for i := 0; i <= idx; i++ {
		res[i] = word[idx-i]
	}
	for i := idx + 1; i < n; i++ {
		res[i] = word[i]
	}

	return string(res)
}

func main() {
	fmt.Println(reversePrefix("abcdefd", 'd'))
	fmt.Println(reversePrefix("xyxzxe", 'z'))
	fmt.Println(reversePrefix("abcd", 'z'))
}
