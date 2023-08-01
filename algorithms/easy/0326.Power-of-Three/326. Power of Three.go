package main

import "fmt"

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}

func main() {
	fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(9))
	fmt.Println(isPowerOfThree(81))
	fmt.Println(isPowerOfThree(121))
}
