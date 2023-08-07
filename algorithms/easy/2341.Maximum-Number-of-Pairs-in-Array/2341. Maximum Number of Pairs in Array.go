package main

import "fmt"

func numberOfPairs(nums []int) []int {
	count := 0

	values := make([]int, 101)
	for _, v := range nums {
		values[v]++
	}

	for _, v := range values {
		count += v / 2
	}

	return []int{count, len(nums) - count*2}
}

func main() {
	fmt.Println(numberOfPairs([]int{1, 3, 2, 1, 3, 2, 2}))
}
