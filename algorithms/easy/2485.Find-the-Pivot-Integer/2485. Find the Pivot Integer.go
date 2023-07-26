package main

import "fmt"

func pivotInteger(n int) int {
	lSum, rSum := 1, (n*(n+1))/2
	if n == 1 {
		return 1
	}
	for i := 2; i <= n; i++ {
		lSum += i
		rSum -= (i - 1)
		if lSum == rSum {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(pivotInteger(8))
	fmt.Println(pivotInteger(1))
	fmt.Println(pivotInteger(4))
}
