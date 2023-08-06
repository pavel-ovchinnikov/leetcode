package main

import "fmt"

func commonFactors(a int, b int) int {
	min := a
	if min > b {
		min = b
	}

	res := 1
	for n := 2; n <= min; n++ {
		if a%n == 0 && b%n == 0 {
			res++
		}
	}

	return res
}

func main() {
	fmt.Println(commonFactors(25, 30))
}
