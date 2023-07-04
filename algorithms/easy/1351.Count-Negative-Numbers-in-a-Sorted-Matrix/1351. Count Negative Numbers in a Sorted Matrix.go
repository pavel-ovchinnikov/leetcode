package leetcode

func countNegatives(grid [][]int) int {
	count := 0
	for _, row := range grid {
		posFirstNeg := getFirstNegativeIndex(row)
		if posFirstNeg != -1 {
			count += len(row) - posFirstNeg
		}
	}
	return count
}

func getFirstNegativeIndex(row []int) int {
	l, r := 0, len(row)-1

	for l <= r {
		mid := (l + r) / 2

		if row[mid] >= 0 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if l == len(row) {
		return -1
	}

	return l
}
