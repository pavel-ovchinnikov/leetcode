package main

import "fmt"

func rearrangeArray(nums []int) []int {
	indexPos, indexNeg := 0, 1
	res := make([]int, len(nums))
	for _, num := range nums {
		if num > 0 {
			res[indexPos] = num
			indexPos += 2
		} else {
			res[indexNeg] = num
			indexNeg += 2
		}
	}
	return res
}

func main() {
	fmt.Println(rearrangeArray([]int{3, 1, -2, -5, 2, -4}))
}
