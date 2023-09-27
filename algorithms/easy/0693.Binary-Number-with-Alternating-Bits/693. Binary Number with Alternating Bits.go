package main

import (
	"fmt"
)

func hasAlternatingBits(n int) bool {
	flag := ((n & 1) == 1)
	n >>= 1

	for n > 0 {
		newFlag := ((n & 1) == 1)
		if newFlag == flag {
			return false
		}
		n >>= 1
		flag = !flag
	}
	return true
}

func main() {
	fmt.Println(hasAlternatingBits(5))
}
