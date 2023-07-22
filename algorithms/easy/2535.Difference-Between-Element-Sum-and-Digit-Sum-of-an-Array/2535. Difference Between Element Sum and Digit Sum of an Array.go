package main

import "fmt"

func differenceOfSum(nums []int) int {
	res := 0
	for _, v := range nums {
		res += v
		for v > 0 {
			res -= v % 10
			v = v / 10
		}
	}

	return res
}

func main() {
	fmt.Println(differenceOfSum([]int{1, 15, 6, 3}))
}
