package main

import "fmt"

func sumOfSquares(nums []int) int {
	res := 0
	n := len(nums)
	for i, num := range nums {
		if n%(i+1) == 0 {
			res += num * num
		}
	}
	return res
}

func main() {
	fmt.Println(sumOfSquares([]int{2, 7, 1, 19, 18, 3}))
}
