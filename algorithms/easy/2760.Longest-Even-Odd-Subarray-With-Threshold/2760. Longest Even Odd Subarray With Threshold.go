package main

import "fmt"

func longestAlternatingSubarray(nums []int, threshold int) int {
	n := len(nums)
	maxi := 0
	i := 0
	j := 0
	flag := 0

	for j < n {
		switch flag {
		case 0:
			if nums[j]%2 == 0 && nums[j] <= threshold {
				i = j
				maxi = max(maxi, j-i+1)
				flag = 1
			}
		case 1:
			x := nums[j-1]
			y := nums[j]
			c := x + y
			if c%2 != 0 && nums[j] <= threshold {
				maxi = max(maxi, j-i+1)
			} else {
				flag = 0
				j--
			}
		}
		j++
	}
	return maxi
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(longestAlternatingSubarray([]int{3, 2, 5, 4}, 5))
	fmt.Println(longestAlternatingSubarray([]int{2, 3, 4, 5}, 4))
	fmt.Println(longestAlternatingSubarray([]int{2, 3, 3, 10}, 18))
}
