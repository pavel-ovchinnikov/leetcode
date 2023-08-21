package main

import (
	"fmt"
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	res := make([][]int, 0, len(arr)/2)

	diff := math.MaxInt
	for i := 0; i < len(arr)-1; i++ {
		newDiff := arr[i+1] - arr[i]

		if newDiff < diff {
			diff = newDiff
			res = res[:0]
		}

		if newDiff == diff {
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}

	return res
}

func main() {
	fmt.Println(minimumAbsDifference([]int{3, 8, -10, 23, 19, -4, -14, 27}))
}
