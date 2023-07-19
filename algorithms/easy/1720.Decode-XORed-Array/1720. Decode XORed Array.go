package main

import "fmt"

func decode(encoded []int, first int) []int {
	res := make([]int, len(encoded)+1)
	res[0] = first
	for i := 0; i < len(res)-1; i++ {
		res[i+1] = res[i] ^ encoded[i]
	}
	return res
}

func main() {
	fmt.Println(decode([]int{1, 2, 3}, 1))
	fmt.Println(decode([]int{6, 2, 7, 3}, 4))
}
