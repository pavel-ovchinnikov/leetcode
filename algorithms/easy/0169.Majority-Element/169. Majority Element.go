package main

import "fmt"

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	count := 0
	num := nums[0]

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			num = nums[i]
		}
		if num == nums[i] {
			count++
		} else {
			count--
		}
	}

	return num
}

func main() {
	fmt.Println(majorityElement([]int{3, 2}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
