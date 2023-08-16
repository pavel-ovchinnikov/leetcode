package main

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	res := make([]int, len(arr1))
	hash := make(map[int]int)
	for i := 0; i < len(arr1); i++ {
		hash[arr1[i]]++
	}

	index := 0
	for _, num := range arr2 {
		count := hash[num]
		for j := 0; j < count; j++ {
			res[index] = num
			index++
		}
		delete(hash, num)
	}
	sortIndex := index
	for num, count := range hash {
		for j := 0; j < count; j++ {
			res[index] = num
			index++
		}
		delete(hash, num)
	}

	sort.Ints(res[sortIndex:])

	return res
}

func main() {
	// fmt.Println(relativeSortArray(
	// 	[]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
	// 	[]int{2, 1, 4, 3, 9, 6}))
	fmt.Println(relativeSortArray(
		[]int{28, 6, 22, 8, 44, 17},
		[]int{22, 28, 8, 6}))
}
