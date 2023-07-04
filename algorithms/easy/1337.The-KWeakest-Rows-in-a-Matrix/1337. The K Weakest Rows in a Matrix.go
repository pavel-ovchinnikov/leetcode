package leetcode

import (
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	type pair struct {
		index    int
		soldiers int
	}
	res := make([]int, k)
	pairs := make([]pair, 0, len(mat))

	for i, v := range mat {
		pairs = append(pairs, pair{i, countSoldiers(v)})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return (pairs[i].soldiers == pairs[j].soldiers &&
			pairs[i].index < pairs[j].index) ||
			pairs[i].soldiers < pairs[j].soldiers
	})

	for i := 0; i < k; i++ {
		res[i] = pairs[i].index
	}
	return res
}

func countSoldiers(row []int) int {
	l, r := 0, len(row)-1
	for l <= r {
		mid := l + (r-l)/2
		if row[mid] == 0 {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
