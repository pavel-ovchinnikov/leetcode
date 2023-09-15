package main

import "fmt"

func countCompleteSubarrays(nums []int) int {
	counter := map[int]int{}
	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++
	}
	uniqueNum := len(counter)
	res := 0

	for i := 0; i < len(nums); i++ {
		hash := map[int]int{}

		for j := i; j < len(nums); j++ {
			hash[nums[j]]++

			if len(hash) == uniqueNum {
				res += len(nums) - j
				break
			}
		}
	}

	return res
}

func main() {
	fmt.Println(countCompleteSubarrays([]int{5, 5, 5, 5}))
}
