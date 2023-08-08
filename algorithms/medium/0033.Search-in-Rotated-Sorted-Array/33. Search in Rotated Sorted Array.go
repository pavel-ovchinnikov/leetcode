package main

import "fmt"

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2
		fmt.Println(nums[l], nums[mid], nums[r])
		if nums[mid] == target {
			return mid
		}

		if nums[l] <= nums[mid] {
			if target >= nums[l] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	if nums[l] == target {
		return l
	}

	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
