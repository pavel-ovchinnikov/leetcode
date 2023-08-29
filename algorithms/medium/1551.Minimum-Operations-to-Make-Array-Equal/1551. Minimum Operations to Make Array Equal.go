package main

import "fmt"

func minOperations(n int) int {
	return n * n / 4
}

func main() {
	fmt.Println(minOperations(3))
	fmt.Println(minOperations(6))
	fmt.Println(minOperations(10))
}
