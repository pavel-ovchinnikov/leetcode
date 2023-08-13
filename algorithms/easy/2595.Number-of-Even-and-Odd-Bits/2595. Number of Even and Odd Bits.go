package main

import "fmt"

func evenOddBit(n int) []int {
	even, odd := 0, 0
	for n > 0 {
		even += n & 1
		n = n >> 1
		odd += n & 1
		n = n >> 1
	}

	return []int{even, odd}
}

func main() {
	fmt.Println(evenOddBit(17))
}
