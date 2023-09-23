package main

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {
	// var res string
	n1, n2 := len(num1), len(num2)
	i1, i2 := n1-1, n2-1
	var num byte

	res := make([]byte, max(n1, n2)+1)
	i := len(res) - 1

	for i1 >= 0 && i2 >= 0 {
		a := num1[i1] - '0'
		b := num2[i2] - '0'
		num += a + b
		res[i] = '0' + num%10
		num /= 10
		i1--
		i2--
		i--
	}

	for i1 >= 0 {
		a := num1[i1] - '0'
		num += a
		res[i] = '0' + num%10
		num /= 10
		i1--
		i--
	}

	for i2 >= 0 {
		b := num2[i2] - '0'
		num += b
		res[i] = '0' + num%10
		num /= 10
		i2--
		i--
	}

	if num == 0 {
		return string(res[1:])
	}

	res[i] = '0' + num
	return string(res)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(addStrings("234", "316"))
	fmt.Println(addStrings("111", "29"))
	fmt.Println(addStrings("29", "111"))
	fmt.Println(addStrings("111", "0"))
	fmt.Println(addStrings("95", "6"))
}
