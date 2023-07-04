package main

import (
	"fmt"
	"sort"
)

func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0, len(nums1))
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})
	idx1, idx2 := 0, 0
	n1, n2 := len(nums1), len(nums2)
	lastVal := -1
	for idx1 < n1 && idx2 < n2 {
		if nums1[idx1] == nums2[idx2] && lastVal != nums1[idx1] {
			res = append(res, nums1[idx1])
			lastVal = nums1[idx1]
			idx1++
			idx2++
			continue
		}

		if nums1[idx1] < nums2[idx2] {
			idx1++
		} else {
			idx2++
		}
	}
	return res
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{1, 2, 2, 1, 5, 6, 7, 7, 7, 7, 3}, []int{7, 2, 2, 3, 5, 67, 122, 12, 3, 4, 5, 6}))
}
