package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {

	left := math.MaxInt
	mid := math.MaxInt
	for _, n := range nums {
		if n <= left {
			left = n
		} else if n <= mid {
			mid = n
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
	fmt.Println(increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
}
