package main

import (
	"fmt"
	"sort"
)

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	hash := map[int]int{}
	for _, pair := range items1 {
		hash[pair[0]] = pair[1]
	}
	for _, pair := range items2 {
		hash[pair[0]] += pair[1]
	}
	res := [][]int{}
	for key, val := range hash {
		res = append(res, []int{key, val})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
	return res
}

// func mergeSimilarItems2(items1 [][]int, items2 [][]int) [][]int {
// 	sort.Slice(items1, func(i, j int) bool {
// 		return items1[i][0] < items1[j][0]
// 	})

// 	sort.Slice(items2, func(i, j int) bool {
// 		return items2[i][0] < items2[j][0]
// 	})

// 	fmt.Println(items1)
// 	fmt.Println(items2)

// 	i1, i2 := 0, 0
// 	n1, n2 := len(items1), len(items2)

// 	res := [][]int{}

// 	for i1 < n1 && i2 < n2 {
// 		switch {
// 		case items1[i1][0] == items2[i2][0]:
// 			res = append(res, []int{items1[i1][0], items1[i1][1] + items2[i2][1]})
// 			i1++
// 			i2++
// 		case items1[i1][0] < items2[i2][0]:
// 			res = append(res, items1[i1])
// 			i1++
// 		case items1[i1][0] > items2[i2][0]:
// 			res = append(res, items2[i2])
// 			i2++
// 		}
// 	}

// 	for i1 < n1 {
// 		res = append(res, items1[i1])
// 		i1++
// 	}

// 	for i2 < n2 {
// 		res = append(res, items2[i2])
// 		i2++
// 	}

// 	return res
// }

func main() {
	items1 := [][]int{
		{1, 1},
		{4, 5},
		{3, 8}}
	items2 := [][]int{
		{3, 1},
		{1, 5},
	}
	fmt.Println(mergeSimilarItems(items1, items2))
}
