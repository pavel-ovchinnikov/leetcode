package main

import "fmt"

func alternateDigitSum(n int) int {
	sum1, sum2 := 0, 0
	digit := 0
	flag := true
	for n > 0 {
		n, digit = n/10, n%10
		if flag {
			sum1 += digit
			sum2 -= digit
		} else {
			sum1 -= digit
			sum2 += digit
		}
		flag = !flag
	}
	if !flag {
		return sum1
	}

	return sum2
}

func main() {
	fmt.Println(alternateDigitSum(10))
}
