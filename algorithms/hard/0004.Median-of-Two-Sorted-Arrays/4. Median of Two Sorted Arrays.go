package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := append(nums1, nums2...)
	sort.Ints(merged)
	n := len(merged)
	if n%2 == 0 {
		return float64(merged[n/2])/2 + float64(merged[n/2-1])/2
	} else {
		return float64(merged[(n-1)/2])
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3}, []int{4, 5, 6}))
}
