package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	isNegative := false
	idx := 0
	res := 0.0

	for idx < len(s) && s[idx] == ' ' {
		idx++
	}

	if idx >= len(s) {
		return 0
	}

	if s[idx] == '-' || s[idx] == '+' {
		isNegative = s[idx] == '-'
		idx++
	}

	for idx < len(s) && s[idx] >= '0' && s[idx] <= '9' {
		res = res*10 + float64(s[idx]) - '0'
		idx++
	}

	if isNegative {
		res *= -1
	}
	res = math.Max(res, float64(math.MinInt32))
	res = math.Min(res, float64(math.MaxInt32))

	return int(res)
}

func main() {
	fmt.Println(myAtoi("  -42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("  +123 "))
	fmt.Println(myAtoi("with words +123"))
	fmt.Println(myAtoi("91283472332"))
	fmt.Println(myAtoi("3.14159"))
	fmt.Println(myAtoi(" .14159"))
	fmt.Println(myAtoi("+-123.1"))
	fmt.Println(myAtoi("-+3.1"))
	fmt.Println(myAtoi("+2.1"))
	fmt.Println(myAtoi("-2.1"))
}
