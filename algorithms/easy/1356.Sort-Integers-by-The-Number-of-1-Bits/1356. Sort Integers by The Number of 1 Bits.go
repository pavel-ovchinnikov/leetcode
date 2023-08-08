package main

import (
	"fmt"
	"math/bits"
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		a, b := bits.OnesCount(uint(arr[i])), bits.OnesCount(uint(arr[j]))
		if a == b {
			return arr[i] < arr[j]
		}
		return a < b
	})
	return arr
}

func main() {

	fmt.Println(sortByBits([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
}
