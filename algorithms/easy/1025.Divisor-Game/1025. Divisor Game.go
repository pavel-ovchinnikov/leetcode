package main

import "fmt"

func divisorGame(n int) bool {
	return (n & 1) == 0
}

func main() {
	fmt.Println(divisorGame(14))
}
