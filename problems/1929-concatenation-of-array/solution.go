package concatenationofarray

func getConcatenation(nums []int) []int {
	res := make([]int, 0)
	res = append(res, nums...)
	res = append(res, nums...)

	return res
}
