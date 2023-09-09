package main

import "fmt"

func permute(nums []int) [][]int {
	res := make([][]int, 0)

	var helper func(idx int, comb []int)
	helper = func(idx int, comb []int) {
		if idx == len(comb) {
			c := make([]int, len(comb))
			copy(c, comb)
			res = append(res, c)
			return
		}
		for i := idx; i < len(comb); i++ {
			comb[idx], comb[i] = comb[i], comb[idx]
			helper(idx+1, comb)
			comb[i], comb[idx] = comb[idx], comb[i]
		}
	}

	helper(0, nums)
	return res
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{1, 2, 3, 4}))
}
