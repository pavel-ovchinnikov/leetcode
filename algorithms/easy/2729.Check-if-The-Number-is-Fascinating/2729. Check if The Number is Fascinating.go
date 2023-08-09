package main

import "fmt"

func isFascinating(n int) bool {

	n2, n3 := 2*n, 3*n
	arr := make([]bool, 10)

	if !(check(n, arr) && check(n2, arr) && check(n3, arr)) {
		return false
	}
	for i := 1; i < len(arr); i++ {
		if !arr[i] {
			return false
		}
	}
	return true
}

func check(n int, arr []bool) bool {
	for d := 0; n > 0; n /= 10 {
		d = n % 10
		switch {
		case d == 0:
			return false
		case arr[d]:
			return false
		default:
			arr[d] = true
		}
	}
	return true
}

func main() {
	fmt.Println(isFascinating(192))
	fmt.Println(isFascinating(100))
}
