package main

import "fmt"

func isHappy(n int) bool {
	count := 0
	for n != 1 && count < 100 {
		n = helper(n)
		count++
	}
	return n == 1
}

func helper(n int) (res int) {
	for n > 0 {
		d := n % 10
		n = n / 10
		res += d * d
	}
	return res
}

func main() {
	fmt.Println(isHappy(19))
}
