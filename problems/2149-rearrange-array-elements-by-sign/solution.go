package rearrangearrayelementsbysign

func rearrangeArray(nums []int) []int {
	res := make([]int, len(nums))
	pos, neg := 0, 1

	for _, n := range nums {
		if n > 0 {
			res[pos] = n
			pos += 2
		} else {
			res[neg] = n
			neg += 2
		}
	}

	return res
}
