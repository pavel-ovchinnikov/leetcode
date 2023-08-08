package main

import "fmt"

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 0; i < len(res); i++ {
		res[i] = res[i>>1] + i&1
	}
	return res
}

func main() {
	fmt.Println(countBits(5))
}
