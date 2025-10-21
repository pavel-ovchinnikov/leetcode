package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	r := 1
	for i := r; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[r] = nums[i]
			r++
		}
	}

	return r
}
