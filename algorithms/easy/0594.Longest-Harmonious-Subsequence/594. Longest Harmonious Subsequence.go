package main

import "fmt"

func findLHS(nums []int) int {
	mm := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mm[nums[i]]++
	}

	res := 0
	for val, count := range mm {
		if count2, ok := mm[val+1]; ok {
			res = Max(res, count+count2)
		}
	}

	return res
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
}
