package main

import "fmt"

func findMatrix(nums []int) [][]int {
	hash := make(map[int]int)
	for _, num := range nums {
		hash[num]++
	}

	res := [][]int{}

	for {
		arr := []int{}
		for num, count := range hash {
			if count > 0 {
				arr = append(arr, num)
				hash[num]--
			}
		}
		if len(arr) == 0 {
			break
		}
		res = append(res, arr)
	}

	return res
}

func main() {
	fmt.Println(findMatrix([]int{1, 3, 4, 1, 2, 3, 1}))
}
