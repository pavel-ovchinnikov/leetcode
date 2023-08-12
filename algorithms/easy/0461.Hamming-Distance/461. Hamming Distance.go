package main

import "fmt"

func hammingDistance(x int, y int) int {
	res := 0
	n := x ^ y
	for n > 0 {
		res++
		n &= n - 1
	}
	return res
}

func main() {
	fmt.Println(hammingDistance(3, 1))
}
