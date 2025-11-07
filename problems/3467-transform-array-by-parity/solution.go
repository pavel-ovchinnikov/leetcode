package transformarraybyparity

func transformArray(nums []int) []int {
	l := len(nums)
	oddCount := 0
	res := make([]int, l)

	for i := 0; i < l; i++ {
		if nums[i]%2 == 1 {
			oddCount++
			res[l-oddCount] = 1
		}
	}

	return res
}
