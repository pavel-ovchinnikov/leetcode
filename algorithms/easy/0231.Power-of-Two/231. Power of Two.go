package main

import "fmt"

func isPowerOfTwo(n int) bool {
	for n > 1 && n%2 == 0 {
		n >>= 1
	}

	return n == 1
}

func main() {
	fmt.Println(isPowerOfTwo(16))
}
