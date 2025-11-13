package minimizemaximumpairsuminarray

import "sort"

func minPairSum(nums []int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	res := 0
	for l < r {
		if sumPair := nums[l] + nums[r]; sumPair > res {
			res = sumPair
		}
		l++
		r--
	}

	return res
}
