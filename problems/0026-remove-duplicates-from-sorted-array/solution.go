package tasks0026

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	i, j := 1, 0
	for ; i < len(nums); i++ {
		if nums[j] != nums[i] {
			nums[j+1] = nums[i]
			j++
		}
	}

	return j + 1
}
