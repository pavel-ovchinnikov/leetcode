package task0217

import "slices"

func containsDuplicate(nums []int) bool {
	hash := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		_, ok := hash[v]
		if ok {
			return true
		}
		hash[v] = struct{}{}
	}

	return false
}

func containsDuplicate_v2(nums []int) bool {
	slices.Sort(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
