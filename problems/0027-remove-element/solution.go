package removeelement

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	k := -1

	for _, v := range nums {
		if v == val {
			continue
		}

		k++
		nums[k] = v
	}

	return k + 1
}
