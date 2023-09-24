package main

import (
	"fmt"
)

func findComplement(num int) int {
	mask := num
	mask |= num >> 1
	mask |= num >> 2
	mask |= num >> 4
	mask |= num >> 8
	mask |= num >> 16
	return num ^ mask
}

func main() {
	fmt.Println(findComplement(5))
}
