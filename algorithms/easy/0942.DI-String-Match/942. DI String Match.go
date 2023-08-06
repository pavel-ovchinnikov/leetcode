package main

import "fmt"

func diStringMatch(s string) []int {
	n := len(s)
	res := make([]int, n+1)
	index := 0
	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			res[i] = index
			index++
		}
	}

	res[n] = index
	index++

	for i := n - 1; i >= 0; i-- {
		if s[i] == 'D' {
			res[i] = index
			index++
		}
	}

	return res
}

func main() {
	fmt.Println(diStringMatch("DDIDDID"))
}
