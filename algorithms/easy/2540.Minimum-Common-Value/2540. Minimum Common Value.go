package main

import (
	"fmt"
	"sort"
)

// func getCommon(nums1 []int, nums2 []int) int {
// 	idx1, idx2 := 0, 0
// 	size1, size2 := len(nums1), len(nums2)
// 	for idx1 < size1 && idx2 < size2 {
// 		if nums1[idx1] == nums2[idx2] {
// 			return nums1[idx1]
// 		}
// 		if nums1[idx1] > nums2[idx2] {
// 			idx2++
// 		} else {
// 			idx1++
// 		}

// 	}
// 	return -1
// }

func getCommon(nums1 []int, nums2 []int) int {
	idx1, idx2 := 0, 0
	size1, size2 := len(nums1), len(nums2)
	for idx1 < size1 && idx2 < size2 {
		idx2 = sort.SearchInts(nums2, nums1[idx1])
		if idx2 >= size2 {
			return -1
		}
		if nums2[idx2] == nums1[idx1] {
			return nums1[idx1]
		}

		idx1 = sort.SearchInts(nums1, nums2[idx2])
		if idx1 >= size1 {
			return -1
		}

		if nums2[idx2] == nums1[idx1] {
			return nums1[idx1]
		}
	}
	return -1
}

func main() {

	fmt.Println(getCommon([]int{1, 2, 3}, []int{2, 4}))
	fmt.Println(getCommon([]int{1, 2, 3, 4}, []int{4, 5, 7}))
	fmt.Println(getCommon([]int{1, 2, 3}, []int{1, 3}))
}
