package findthearrayconcatenationvalue

import "strconv"

func findTheArrayConcVal(nums []int) int64 {
	if len(nums) == 0 {
		return 0
	}

	p1, p2 := 0, len(nums)-1
	var res int64 = 0

	for p1 < p2 {
		l, r := strconv.Itoa(nums[p1]), strconv.Itoa(nums[p2])
		n, _ := strconv.ParseInt(l+r, 10, 64)
		res += int64(n)

		p1++
		p2--
	}

	if p1 == p2 {
		res += int64(nums[p1])
	}

	return res
}
