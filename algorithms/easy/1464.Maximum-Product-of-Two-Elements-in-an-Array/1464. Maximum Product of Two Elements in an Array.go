package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	max1, max2 := math.MinInt, math.MinInt
	for _, num := range nums {
		if num >= max1 {
			max2 = max1
			max1 = num
		} else if num >= max2 {
			max2 = num
		}
	}

	return (max1 - 1) * (max2 - 1)
}

func main() {
	fmt.Println(maxProduct([]int{3, 4, 5, 2}))
	fmt.Println(maxProduct([]int{1, 5, 4, 5}))
	fmt.Println(maxProduct([]int{3, 7}))
	fmt.Println(maxProduct([]int{10, 2, 5, 2}))
}
