package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	return x == revert(x)
}

func revert(x int) int {
	a := 0
	for x != 0 {
		a = a*10 + x%10
		x = x / 10
	}
	return a
}

func main() {
	fmt.Println(isPalindrome(100))
	fmt.Println(isPalindrome(232))
}
