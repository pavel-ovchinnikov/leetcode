package main

import "fmt"

func createTargetArray(nums []int, index []int) []int {
	res := make([]int, len(nums))

	for idx, i := range index {
		for j := len(res) - 1; j > i; j-- {
			res[j] = res[j-1]
		}

		res[i] = nums[idx]
	}

	return res
}

func main() {
	fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
}
