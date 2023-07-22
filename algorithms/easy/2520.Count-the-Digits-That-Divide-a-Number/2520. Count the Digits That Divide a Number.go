package main

import "fmt"

func countDigits(num int) int {
	res := 0
	for i := num; i > 0; i /= 10 {
		if (num % (i % 10)) == 0 {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(countDigits(7))
	fmt.Println(countDigits(121))
	fmt.Println(countDigits(1248))
}
