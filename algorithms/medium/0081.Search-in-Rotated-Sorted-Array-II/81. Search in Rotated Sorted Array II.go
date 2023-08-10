package main

import "fmt"

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + (r-l)/2

		if nums[mid] == target {
			return true
		}

		if nums[l] == nums[mid] && nums[mid] == nums[r] {
			l++
			r--
			continue
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
	return nums[l] == target
}

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	fmt.Println(search([]int{1, 0, 1, 1, 1}, 0))
}
