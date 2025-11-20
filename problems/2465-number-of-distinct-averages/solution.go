package numberofdistinctaverages

import "sort"

func distinctAverages(nums []int) int {
	i, j := 0, len(nums)-1
	seen := make(map[float64]struct{})

	sort.Ints(nums)
	for i < j {
		avg := float64(nums[i]+nums[j]) / 2.0
		seen[avg] = struct{}{}
		i++
		j--
	}

	return len(seen)
}
