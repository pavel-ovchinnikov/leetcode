package main

import "fmt"

func arrangeCoins(n int) int {
	rows := 0
	col := 1
	for n >= col {
		n -= col
		col++
		rows++
	}
	return rows
}

func main() {
	fmt.Println(arrangeCoins(1))
	fmt.Println(arrangeCoins(3))
	fmt.Println(arrangeCoins(6))
	fmt.Println(arrangeCoins(10))
	fmt.Println(arrangeCoins(123))
}
