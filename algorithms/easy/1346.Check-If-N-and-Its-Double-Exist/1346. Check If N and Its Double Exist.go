package main

import (
	"fmt"
	"sort"
)

func checkIfExist(arr []int) bool {
	sort.Ints(arr)
	for idx, val := range arr {
		i := sort.SearchInts(arr, val*2)
		if idx != i && i < len(arr) && arr[i] == val*2 {
			fmt.Println(val, arr[i])
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkIfExist([]int{10, 2, 5, 3}))
	fmt.Println(checkIfExist([]int{1, 2, 5, 4}))
	// [-2,0,10,-19,4,6,-8]
	fmt.Println(checkIfExist([]int{-2, 0, 10, -19, 4, 6, -8}))
}
