package main

import "fmt"

// func isSameAfterReversals(num int) bool {
//     return num == 0 || num % 10 > 0
// }

func isSameAfterReversals(num int) bool {
	return revers(revers(num)) == num
}

func revers(a int) (res int) {
	b := 0
	for a > 0 {
		a, b = a/10, a%10
		res = res*10 + b
	}
	return res
}

func main() {
	fmt.Println(isSameAfterReversals(542))
	fmt.Println(isSameAfterReversals(242))
	fmt.Println(isSameAfterReversals(2420))
}
