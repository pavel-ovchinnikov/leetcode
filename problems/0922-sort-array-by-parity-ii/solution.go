package sortarraybyparityii

func sortArrayByParityII(nums []int) []int {
	i, j := 0, 1

	for i < len(nums) && j < len(nums) {
		if nums[i]%2 == 0 {
			i += 2
			continue
		}

		if nums[j]%2 == 1 {
			j += 2
			continue
		}

		nums[i], nums[j] = nums[j], nums[i]
		i += 2
		j += 2
	}

	return nums
}
