package main

import (
	"fmt"
	"sort"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr2)

	res := 0
	for _, val := range arr1 {
		minIdx := sort.SearchInts(arr2, val-d)
		if minIdx < len(arr2) && arr2[minIdx] >= val-d && arr2[minIdx] <= val+d {
			continue
		}
		res++
	}

	return res
}

func main() {
	fmt.Println(findTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2))
}
