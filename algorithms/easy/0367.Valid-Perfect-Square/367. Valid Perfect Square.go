package main

import (
	"fmt"
	"math"
)

func isPerfectSquare(num int) bool {
	val := int(math.Sqrt(float64(num)))
	return val*val == num
}

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
}
