package main

import "fmt"

func minimumReplacement(nums []int) int64 {
	n := len(nums)
	res := int64(0)
	last := nums[n-1]

	for i := n - 2; i >= 0; i-- {
		if nums[i] > last {
			split := nums[i] / last
			if nums[i]%last != 0 {
				split++
				last = nums[i] / split
			}
			res += int64(split - 1)
			continue
		}

		last = nums[i]
	}

	return res
}

func main() {
	fmt.Println(minimumReplacement([]int{3, 9, 3}))
	fmt.Println(minimumReplacement([]int{1, 2, 3, 4, 5}))
}
