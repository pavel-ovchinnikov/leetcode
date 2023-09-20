package main

import "fmt"

func isPowerOfFour(n int) bool {
	if n < 1 {
		return false
	}

	if n == 1 {
		return true
	}

	for n%4 == 0 {
		n /= 4
	}

	return n == 1
}

func main() {
	fmt.Println(isPowerOfFour(0))
	fmt.Println(isPowerOfFour(-16))
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(256))
	fmt.Println(isPowerOfFour(512))
	fmt.Println(isPowerOfFour(8))
}
