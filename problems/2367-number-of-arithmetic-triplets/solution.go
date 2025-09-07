package numberofarithmetictriplets

func arithmeticTriplets(nums []int, diff int) int {
	res := 0
	i, j, k := 0, 0, 0

	for i = 0; i < len(nums); i++ {
		for j = i + 1; j < len(nums); j++ {
			if nums[j]-nums[i] == diff {
				for k = j + 1; k < len(nums); k++ {
					if nums[k]-nums[j] == diff {
						res++
						break
					}
				}
			}
		}
	}

	return res
}
