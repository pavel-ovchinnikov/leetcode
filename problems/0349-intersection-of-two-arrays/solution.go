package intersectionoftwoarrays

import "sort"

func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	i, j := 0, 0

	result := make([]int, 0, min(len(nums1), len(nums2)))
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
			continue
		}

		if nums1[i] > nums2[j] {
			j++
			continue
		}

		if len(result) > 0 && result[len(result)-1] == nums1[i] {
			i++
			j++
			continue
		}

		result = append(result, nums1[i])
		i++
		j++

	}

	return result
}
