package minimumaverageofsmallestandlargestelements

import (
	"math"
	"sort"
)

func minimumAverage(nums []int) float64 {
	minAVG := math.MaxFloat64
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	for l < r {
		avg := float64(nums[l]+nums[r]) / 2
		if avg < minAVG {
			minAVG = avg
		}
		l++
		r--
	}

	return minAVG
}
