package main

import (
	"fmt"
	"strconv"
)

func countSymmetricIntegers(low int, high int) int {
	res := 0
	for i := low; i <= high; i++ {
		if !isSymmetric(i) {
			continue
		}
		res++
	}
	return res
}

func isSymmetric(num int) bool {
	s := strconv.Itoa(num)
	n := len(s)
	if n%2 != 0 {
		return false
	}

	sum := 0
	for i := 0; i < len(s)/2; i++ {
		sum += int(s[i] - '0')
		sum -= int(s[n-i-1] - '0')
	}
	return sum == 0
}

func main() {
	fmt.Println(countSymmetricIntegers(1, 101))
}
