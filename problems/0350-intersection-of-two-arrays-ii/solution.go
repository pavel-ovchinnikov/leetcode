package task0350

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	res := make([]int, 0, min(len(nums1), len(nums2)))
	sort.Ints(nums1)
	sort.Ints(nums2)

	i1, i2 := 0, 0
	for i1 < len(nums1) && i2 < len(nums2) {
		if nums1[i1] == nums2[i2] {
			res = append(res, nums1[i1])
			i1++
			i2++
			continue
		}

		if nums1[i1] < nums2[i2] {
			i1++
		} else {
			i2++
		}
	}

	return res
}
