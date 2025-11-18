package findallkdistantindicesinanarray

import "sort"

func findKDistantIndices(nums []int, key int, k int) []int {
	n := len(nums)
	indicesMap := make(map[int]struct{})
	for i := 0; i < n; i++ {
		if nums[i] == key {
			start := i - k
			if start < 0 {
				start = 0
			}
			end := i + k
			if end >= n {
				end = n - 1
			}
			for j := start; j <= end; j++ {
				indicesMap[j] = struct{}{}
			}
		}
	}

	nums = nums[:0]
	for index := range indicesMap {
		nums = append(nums, index)
	}
	sort.Ints(nums)
	return nums
}
