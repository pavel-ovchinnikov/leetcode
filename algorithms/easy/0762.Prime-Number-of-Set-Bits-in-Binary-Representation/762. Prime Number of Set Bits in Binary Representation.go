package main

import (
	"fmt"
	"math/bits"
)

func countPrimeSetBits(left int, right int) int {
	res := 0

	for i := left; i <= right; i++ {
		if isPrime(bits.OnesCount(uint(i))) {
			res++
		}
	}

	return res
}

func isPrime(num int) bool {
	switch {
	case num <= 1:
		return false
	case num <= 3:
		return true
	case num%2 == 0 || num%3 == 0:
		return false
	}

	for k := 5; k*k < num; k += 6 {
		if num%k == 0 || num%(k+2) == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(countPrimeSetBits(6, 10))
}
