package main

import "fmt"

func arraySign(nums []int) int {
	isPositive := true

	for _, num := range nums {
		if num == 0 {
			return 0
		}

		if num < 0 {
			isPositive = !isPositive
		}
	}

	if isPositive {
		return 1
	} else {
		return -1
	}
}

func main() {
	fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1}))
	fmt.Println(arraySign([]int{1, 5, 0, 2, -3}))
	fmt.Println(arraySign([]int{-1, 1, -1, 1, -1}))
}
