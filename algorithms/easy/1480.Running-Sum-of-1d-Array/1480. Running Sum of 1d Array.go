package main

import "fmt"

func runningSum(nums []int) []int {
	sum := 0

	for index, num := range nums {
		sum += num
		nums[index] = sum
	}

	return nums
}

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))
	fmt.Println(runningSum([]int{3, 1, 2, 10, 1}))
}
