package main

import "fmt"

func singleNumber(nums []int) int {
	result := 0
	for _, v := range nums {
		result = result ^ v
	}
	return result
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
}
