package main

import "fmt"

func digitCount(num string) bool {
	arr := make([]int, 10)
	for _, c := range num {
		arr[int(c-'0')]++
	}

	for i, c := range num {
		if arr[i] != int(c-'0') {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(digitCount("1210"))
}
