package main

import "fmt"

func numberOfSubarrays(nums []int, k int) int {
	n := len(nums)
	res, odd, nice := 0, 0, 0

	for l, r := 0, 0; r < n; r++ {
		if nums[r]%2 != 0 {
			odd++
			nice = 0
		}

		for odd == k {
			nice++
			if nums[l]%2 != 0 {
				odd--
			}
			l++
		}

		res += nice
	}

	return res
}

func main() {
	fmt.Println(numberOfSubarrays([]int{1, 1, 2, 1, 1}, 3))                // 2
	fmt.Println(numberOfSubarrays([]int{2, 4, 6}, 1))                      // 0
	fmt.Println(numberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2)) // 16
}
