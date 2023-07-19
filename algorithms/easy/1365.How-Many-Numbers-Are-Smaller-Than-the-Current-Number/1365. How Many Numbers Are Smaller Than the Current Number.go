package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))
	var count = 0
	for i := 0; i < len(nums); i++ {
		for _, val := range nums {
			if val != nums[i] && val < nums[i] {
				count++
			}
		}
		res[i] = count
		count = 0
	}
	return res
}

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
	fmt.Println(smallerNumbersThanCurrent([]int{6, 5, 4, 8}))
}
