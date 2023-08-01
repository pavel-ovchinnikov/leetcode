package main

import "fmt"

func convertToTitle(columnNumber int) string {
	var res string

	for columnNumber != 0 {
		index := (columnNumber - 1) % 26
		columnNumber = (columnNumber - 1) / 26
		res = string(rune(index)+65) + res
	}

	return res
}

func main() {
	fmt.Println(convertToTitle(28))
}
