package main

import "fmt"

func titleToNumber(columnTitle string) int {
	res := 0
	for _, c := range columnTitle {
		res *= 26
		res += int(c) - 'A' + 1
	}

	return res
}

func main() {
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("Z"))
	fmt.Println(titleToNumber("AZ"))
	fmt.Println(titleToNumber("ZY"))

}
