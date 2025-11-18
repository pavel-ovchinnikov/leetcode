package largestpositiveintegerthatexistswithitsnegative

import "sort"

func findMaxK(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1

	for l < r {
		if nums[l] >= 0 {
			return -1
		}
		if nums[r] <= 0 {
			return -1
		}

		if -1*nums[l] == nums[r] {
			return nums[r]
		}

		if -1*nums[l] > nums[r] {
			l++
		}

		if -1*nums[l] < nums[r] {
			r--
		}
	}

	return -1
}
