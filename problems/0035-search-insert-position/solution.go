package tasks0035

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	l, r := 0, len(nums)-1

	for l < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if nums[l] < target {
		return l + 1
	}

	return l
}
