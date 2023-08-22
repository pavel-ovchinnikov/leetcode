package main

import "fmt"

func intersection(nums [][]int) []int {
	n := len(nums)
	mem := make([]int, 1001)
	res := make([]int, 0, len(nums[0]))

	for _, row := range nums {
		for _, num := range row {
			mem[num]++
		}
	}

	for i, num := range mem {
		if num == n {
			res = append(res, i)
		}
	}

	return res
}

func main() {
	fmt.Println(intersection([][]int{
		{3, 1, 2, 4, 5},
		{1, 2, 3, 4},
		{3, 4, 5, 6},
	}))
}
