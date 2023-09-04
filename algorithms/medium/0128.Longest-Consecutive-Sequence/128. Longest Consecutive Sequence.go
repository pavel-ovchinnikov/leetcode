package main

import (
	"fmt"
)

// func longestConsecutive(nums []int) int {
// 	hash := map[int]int{} // [num]length
// 	maxCon := 0
// 	updateHash := func(num int) int {
// 		hash[num] = hash[num-1] + 1
// 		num++
// 		for hash[num] != 0 {
// 			hash[num] = hash[num-1] + 1
// 			num++
// 		}
// 		return hash[num-1]
// 	}

// 	for _, num := range nums {
// 		if hash[num] != 0 {
// 			continue
// 		}

// 		maxCon = max(maxCon, updateHash(num))
// 	}
// 	return maxCon
// }

func longestConsecutive(nums []int) int {
	numMap := map[int]bool{}
	countCons := 0

	for _, num := range nums {
		numMap[num] = true
	}

	for num := range numMap {
		if numMap[num-1] {
			continue
		}

		cur := num
		for numMap[cur+1] {
			cur++
		}

		countCons = max(countCons, cur-num+1)
	}

	return countCons
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {

	fmt.Println(longestConsecutive([]int{1, 2, 0, 1}))
	// fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	// fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
