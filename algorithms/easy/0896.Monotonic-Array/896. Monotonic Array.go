package main

import "fmt"

func isMonotonic(nums []int) bool {
	dir := 0 // direction: -1, 0, -1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			if dir == 0 {
				dir = 1
			}

			if dir != 1 {
				return false
			}
			continue
		}

		if nums[i] < nums[i-1] {
			if dir == 0 {
				dir = -1
			}

			if dir != -1 {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isMonotonic([]int{1, 2, 3, 4, 5}))
	fmt.Println(isMonotonic([]int{5, 4, 3, 2, 1}))
	fmt.Println(isMonotonic([]int{5, 5, 5, 4, 3, 2, 1}))
	fmt.Println(isMonotonic([]int{5, 5, 5, 6}))
	fmt.Println(isMonotonic([]int{5, 5, 5, 1}))
	fmt.Println(isMonotonic([]int{5, 5, 5, 1, 1}))
	fmt.Println(isMonotonic([]int{5, 5, 5, 7, 7}))

	fmt.Println(isMonotonic([]int{1, 3, 2}))
}
