package main

import "fmt"

func sumBase(n int, k int) int {
	res := 0
	for n > 0 {
		res += n % k
		n /= k
	}
	return res
}

func main() {
	fmt.Println(sumBase(32, 6))
}
