package main

import "fmt"

func smallestEvenMultiple(n int) int {
	if n&1 == 1 {
		return n << 1
	}
	return n
}

func main() {
	fmt.Println(smallestEvenMultiple(1))
	fmt.Println(smallestEvenMultiple(5))
	fmt.Println(smallestEvenMultiple(6))
	fmt.Println(smallestEvenMultiple(16))
}
