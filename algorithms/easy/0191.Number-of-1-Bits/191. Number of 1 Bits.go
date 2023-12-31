package main

import "fmt"

func hammingWeight(num uint32) int {
	if num == 0 {
		return 0
	}

	res := 0
	for num != 0 {
		if num&1 == 1 {
			res++
		}
		num >>= 1
	}
	return res
}

func main() {
	fmt.Println(hammingWeight(11))
}
