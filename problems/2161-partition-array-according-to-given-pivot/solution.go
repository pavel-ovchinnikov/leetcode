package partitionarrayaccordingtogivenpivot

func pivotArray(nums []int, pivot int) []int {
	res := make([]int, len(nums))

	l, r := 0, len(nums)-1
	for i := 0; i < len(nums); i++ {

		if nums[i] < pivot {
			res[l] = nums[i]
			l++
		}

		ii := len(nums) - 1 - i
		if nums[ii] > pivot {
			res[r] = nums[ii]
			r--
		}
	}

	for l <= r {
		res[l] = pivot
		l++
	}

	return res
}
