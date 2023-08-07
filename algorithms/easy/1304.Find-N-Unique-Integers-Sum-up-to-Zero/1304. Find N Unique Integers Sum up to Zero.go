package main

import "fmt"

func sumZero(n int) []int {
	res := make([]int, n)

	for i := 0; i < n/2*2; i += 2 {
		res[i] = -(i + 1)
		res[i+1] = i + 1
	}

	return res
}

func main() {
	fmt.Println(sumZero(1))
	fmt.Println(sumZero(3))
	fmt.Println(sumZero(10))
}
