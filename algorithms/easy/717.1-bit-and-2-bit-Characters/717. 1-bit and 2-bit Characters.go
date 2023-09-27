package main

import "fmt"

func isOneBitCharacter(bits []int) bool {
	count := 0
	for i := len(bits) - 2; i >= 0 && bits[i] != 0; i-- {
		count++
	}
	return (count % 2) == 0
}

func main() {
	fmt.Println(isOneBitCharacter([]int{1, 1, 1, 0}))
}
