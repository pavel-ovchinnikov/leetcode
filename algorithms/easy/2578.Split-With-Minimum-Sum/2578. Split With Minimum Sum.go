package main

import (
	"fmt"
	"sort"
)

func splitNum(num int) int {
	nums := make([]int, 0, 9)

	for num > 0 {
		if num%10 != 0 {
			nums = append(nums, num%10)
		}
		num /= 10
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	factor := 1
	value := 0

	for i := 0; i < len(nums); i += 2 {
		temp := nums[i]
		if i+1 < len(nums) {
			temp += nums[i+1]
		}
		value += temp * factor
		factor *= 10
	}
	return value
}

func main() {
	fmt.Println(splitNum(4325))
}
