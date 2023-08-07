package main

import "fmt"

func selfDividingNumbers(left int, right int) []int {
	res := []int{}
	for i := left; i <= right; i++ {
		if isSelfDiveded(i) {
			res = append(res, i)
		}
	}
	return res
}

func isSelfDiveded(num int) bool {
	for n := num; n > 0; n /= 10 {
		if n%10 == 0 || num%(n%10) != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}
