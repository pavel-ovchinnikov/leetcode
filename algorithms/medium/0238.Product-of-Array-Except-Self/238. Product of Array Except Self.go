package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	l, r := 1, 1
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}

	for i := 0; i < n; i++ {
		res[i] *= l
		l *= nums[i]
	}
	// fmt.Println(res)

	for i := n - 1; i >= 0; i-- {
		res[i] *= r
		r *= nums[i]
	}
	// fmt.Println(res)

	return res
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
