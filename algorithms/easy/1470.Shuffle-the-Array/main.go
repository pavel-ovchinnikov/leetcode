package main

import "fmt"

func shuffle(nums []int, n int) []int {
	result := make([]int, 2*n)
	for i := 0; i < n; i++ {
		result[2*i] = nums[i]
		result[2*i+1] = nums[n+i]
	}
	return result
}

func main() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
}
