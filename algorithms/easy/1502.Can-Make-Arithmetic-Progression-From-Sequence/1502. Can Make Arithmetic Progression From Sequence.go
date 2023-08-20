package main

import (
	"fmt"
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	if len(arr) < 2 {
		return true
	}
	diff := arr[1] - arr[0]

	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != diff {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canMakeArithmeticProgression([]int{3, 5, 1}))
}
