package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int)
	stack := []int{}

	for _, num := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < num {
			val := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			hash[val] = num
		}
		stack = append(stack, num)
	}

	res := make([]int, len(nums1))
	for i, num := range nums1 {
		if val, ok := hash[num]; ok {
			res[i] = val
		} else {
			res[i] = -1
		}
	}
	return res
}

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}
