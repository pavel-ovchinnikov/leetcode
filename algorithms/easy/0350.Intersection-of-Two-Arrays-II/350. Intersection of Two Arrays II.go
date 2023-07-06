package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}
	size1, size2 := len(nums1), len(nums2)
	idx1, idx2 := 0, 0

	sort.Ints(nums1)
	sort.Ints(nums2)
	for idx1 < size1 && idx2 < size2 {
		if nums1[idx1] == nums2[idx2] {
			res = append(res, nums1[idx1])
			idx1++
			idx2++
			continue
		}
		if nums1[idx1] > nums2[idx2] {
			idx2++
		} else {
			idx1++
		}
	}
	return res
}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	fmt.Println(intersect([]int{0, 1}, []int{0, 1, 1, 0, 0}))
}
