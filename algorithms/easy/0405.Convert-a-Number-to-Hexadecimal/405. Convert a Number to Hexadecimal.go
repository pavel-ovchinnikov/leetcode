package main

import "fmt"

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	hex := "0123456789abcdef"

	res := ""
	count := 0
	for num != 0 && count < 8 {
		res = string(hex[num&15]) + res
		num >>= 4
		count++
	}

	return res
}

func main() {
	fmt.Println(toHex(13))
	fmt.Println(toHex(-1))
}
