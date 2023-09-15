package main

import "fmt"

func getAverages(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, n)
	windowSize := 2*k + 1
	windowSum := 0

	for l, r := 0, 0; r < n; r++ {
		windowSum += nums[r]
		res[r] = -1

		if r-l+1 != windowSize {
			continue
		}

		res[r-k] = windowSum / windowSize
		windowSum -= nums[l]
		l++
	}

	return res
}

func main() {
	fmt.Println(getAverages([]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3))
	fmt.Println(getAverages([]int{100000}, 0))
	fmt.Println(getAverages([]int{8}, 1000))

}
