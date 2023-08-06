package main

import "fmt"

func removeTrailingZeros(num string) string {
	index := len(num) - 1
	for index >= 0 && num[index] == '0' {
		index--
	}

	return string([]byte(num[:index+1]))
}

func main() {
	fmt.Println(removeTrailingZeros("12300"))
}
