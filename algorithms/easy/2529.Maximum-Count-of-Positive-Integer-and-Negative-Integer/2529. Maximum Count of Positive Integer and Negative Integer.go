package leetcode

func maximumCount(nums []int) int {
	negCount := search(nums, func(val int) bool {
		return val < 0
	})
	posCount := len(nums) - search(nums, func(val int) bool {
		return val <= 0
	})

	if posCount > negCount {
		return posCount
	}
	return negCount
}

func search(nums []int, checkFn func(val int) bool) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if checkFn(nums[mid]) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
