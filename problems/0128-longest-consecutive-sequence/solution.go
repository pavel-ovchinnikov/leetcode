package task0128

func longestConsecutive(nums []int) int {
	set := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		set[v] = struct{}{}
	}

	res := 0

	for n := range set {
		// Find the start of a consecutive sequence.
		if _, ok := set[n-1]; ok {
			continue
		}

		// Count consecutive numbers starting from n.
		l := 0
		for {
			if _, ok := set[n+l]; !ok {
				break
			}
			l++
		}

		res = max(res, l)
	}

	return res
}
