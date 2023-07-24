package main

import "fmt"

func makeSmallestPalindrome(s string) string {
	res := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		l, r := res[i], res[len(s)-1-i]
		if l != r {
			if l < r {
				res[len(s)-1-i] = l
			} else {
				res[i] = r
			}
		}
	}
	return string(res)
}

func main() {
	fmt.Println(makeSmallestPalindrome("egcfe"))
	fmt.Println(makeSmallestPalindrome("abcd"))
	fmt.Println(makeSmallestPalindrome("seven"))
}
