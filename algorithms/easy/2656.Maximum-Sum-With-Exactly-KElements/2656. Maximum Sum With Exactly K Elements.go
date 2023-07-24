package main

import "fmt"

func maximizeSum(nums []int, k int) int {
	sum := 0
	for _, num := range nums {
		if num > sum {
			sum = num
		}
	}
	start := sum + 1
	last := sum + k
	for start < last {
		sum += start
		start++
	}
	return sum
}

func main() {
	fmt.Println(maximizeSum([]int{1, 2, 3, 4, 5}, 3))
}
