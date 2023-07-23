package main

import "fmt"

func arithmeticTriplets(nums []int, diff int) int {
	res := 0
	hashmap := make(map[int]bool)

	for _, v := range nums {
		hashmap[v] = true
	}
	for i := 0; i < len(nums); i++ {
		if hashmap[nums[i]-diff] && hashmap[nums[i]+diff] {
			res++
		}

	}
	return res
}

func main() {
	fmt.Println(arithmeticTriplets([]int{0, 1, 4, 6, 7, 10}, 3))
	fmt.Println(arithmeticTriplets([]int{4, 5, 6, 7, 8, 9}, 2))
}
