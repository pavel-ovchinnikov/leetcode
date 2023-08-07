package main

import (
	"fmt"
)

func sortString(s string) string {
	n := len(s)
	iRes := 0
	res := make([]byte, n)
	letters := make([]int, 26)

	for _, c := range s {
		letters[c-'a']++
	}

	for iRes < n {
		for i, c := range letters {
			if c != 0 {
				res[iRes] = byte(i + 'a')
				letters[i]--
				iRes++
			}
		}

		for i := len(letters) - 1; i >= 0; i-- {
			c := letters[i]
			if c != 0 {
				res[iRes] = byte(i + 'a')
				letters[i]--
				iRes++
			}
		}
	}

	return string(res)
}

func main() {
	fmt.Println(sortString("aaaabbbbcccc"), sortString("aaaabbbbcccc") == "abccbaabccba")
}
