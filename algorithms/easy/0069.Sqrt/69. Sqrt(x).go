package main

import "fmt"

func mySqrt(x int) int {
	l, r := 0, x/2
	for l <= r {
		mid := l + (r-l)/2

		if mid*mid == x {
			return mid
		}

		if mid*mid > x {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	if l*l > x {
		return l - 1
	}

	return l
}

func main() {
	fmt.Println(mySqrt(8))
}
