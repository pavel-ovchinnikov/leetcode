package leetcode

func countNegatives(grid [][]int) int {
	count := 0
	size := len(grid)

	for i := 0; i < size; i++ {
		row := grid[i]
		l, r := 0, size-1
		posFirstNeg := -1

		for l <= r {
			mid := l + (r-l)/2
			if row[mid] < 0 {
				posFirstNeg = mid
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		if posFirstNeg != -1 {
			count += size - posFirstNeg
		}
	}
	return count
}
