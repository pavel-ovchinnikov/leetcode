package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for index, num := range nums {
		num2 := target - num
		if index2, ok := hashmap[num]; ok {
			return []int{index2, index}
		}
		hashmap[num2] = index
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 3))
}
