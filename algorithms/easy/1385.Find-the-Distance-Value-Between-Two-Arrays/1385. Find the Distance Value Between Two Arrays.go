package main

import (
	"fmt"
	"math"
	"sort"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	sort.Ints(arr1)
	sort.Ints(arr2)

	for _, val1 := range arr1 {
		f := true
		for _, val2 := range arr2 {
			if math.Abs(float64(val1-val2)) <= float64(d) {
				f = false
			}
		}
		if f {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(findTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2))
}
