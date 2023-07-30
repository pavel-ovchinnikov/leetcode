package main

import "fmt"

func subsetXORSum(nums []int) int {
	sum, sumset := 0, 0
	sum, _ = solve(nums, sum, sumset, 0)
	return sum
}

func solve(nums []int, sum, sumset, start int) (int, int) {
	sum += sumset
	for i := start; i < len(nums); i++ {
		sumset ^= nums[i]
		sum, sumset = solve(nums, sum, sumset, i+1)
		sumset ^= nums[i]
	}
	return sum, sumset
}

func main() {
	fmt.Println(subsetXORSum([]int{1, 3}))
	fmt.Println(subsetXORSum([]int{5, 1, 6}))
}
