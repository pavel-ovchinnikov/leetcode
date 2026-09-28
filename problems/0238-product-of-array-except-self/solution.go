package task0238

func productExceptSelf(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	result := make([]int, n)

	product := 1
	for i := 0; i < n; i++ {
		left[i] = product
		product *= nums[i]
	}

	product = 1
	for i := n - 1; i >= 0; i-- {
		right[i] = product
		product *= nums[i]
	}

	for i := 0; i < n; i++ {
		result[i] = left[i] * right[i]
	}
	return result
}
