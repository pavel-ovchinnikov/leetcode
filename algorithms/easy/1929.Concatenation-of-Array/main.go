package main

import "fmt"

func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, n*2)
	for i := 0; i < n*2; i++ {
		ans[i] = nums[i%n]
	}
	return ans
}

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 1}))
}
