package main

import "fmt"

func convertToTitle(columnNumber int) string {
	res := make([]rune, 0, columnNumber/26)

	for columnNumber != 0 {
		index := (columnNumber - 1) % 26
		columnNumber = (columnNumber - 1) / 26
		res = append(res, rune(index)+65)
	}

	reverse(res)
	return string(res)
}

func reverse(word []rune) {
	l, r := 0, len(word)-1
	for l < r {
		word[l], word[r] = word[r], word[l]
		l++
		r--
	}
}

func main() {
	fmt.Println(convertToTitle(28))
}
