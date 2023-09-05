package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	if n < 0 {
		n *= -1
		x = 1 / x
	}

	if n%2 == 0 {
		x = myPow(x*x, n/2)
	} else {
		x = x * myPow(x*x, n/2)
	}
	return x
}

func main() {
	fmt.Println(myPow(2.0, 10))
	fmt.Println(myPow(2.0, 11))
	fmt.Println(myPow(0.5, 2))
	fmt.Println(myPow(2, -2))

}
