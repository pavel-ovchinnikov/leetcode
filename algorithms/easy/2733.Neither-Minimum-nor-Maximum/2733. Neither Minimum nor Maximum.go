package main

import (
	"fmt"
	"math"
)

func findNonMinOrMax(nums []int) int {
	if len(nums) <= 2 {
		return -1
	}

	minNum, maxNum := math.MaxInt, math.MinInt

	for _, num := range nums {
		if num < minNum {
			minNum = num
		}

		if num > maxNum {
			maxNum = num
		}
	}

	for _, num := range nums {
		if minNum < num && num < maxNum {
			return num
		}
	}
	return -1
}

func main() {

	fmt.Println(findNonMinOrMax([]int{10, 27, 3}))
	fmt.Println(findNonMinOrMax([]int{1, 2, 3, 4, 5}))
	fmt.Println(findNonMinOrMax([]int{3, 30, 24}))
	fmt.Println(findNonMinOrMax([]int{2, 1, 3}))

	fmt.Println(findNonMinOrMax([]int{1, 2}))
	fmt.Println(findNonMinOrMax([]int{1, 1, 2, 2}))
	fmt.Println(findNonMinOrMax([]int{1, 1}))
}
