package main

import "fmt"

func numOfSubarrays(arr []int, k int, threshold int) int {
	n := len(arr)
	windowSum := 0
	res := 0

	for l, r := 0, 0; r < n; r++ {
		windowSum += arr[r]

		if r-l+1 != k {
			continue
		}

		if windowSum/k >= threshold {
			res++
		}
		windowSum -= arr[l]
		l++
	}

	return res
}

func main() {
	fmt.Println(numOfSubarrays([]int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}, 3, 5))
}
