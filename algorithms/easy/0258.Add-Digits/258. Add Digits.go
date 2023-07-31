package main

import "fmt"

func addDigits(num int) int {
	return 1 + (num-1)%9
}

func main() {
	fmt.Println(addDigits(38))
}
