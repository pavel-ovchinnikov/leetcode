package problem

func targetIndices(nums []int, target int) []int {
	count, left := 0, 0
	for _, val := range nums {
		if target == val {
			count++
		} else if target > val {
			left++
		}
	}

	res := make([]int, count)
	for i := 0; i < count; i++ {
		res[i] = left
		left++
	}

	return res
}
