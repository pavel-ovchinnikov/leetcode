package missingnumber

func missingNumber(nums []int) int {
	var n int = len(nums)

	sum := (n * (n + 1)) / 2

	for i := 0; i < n; i++ {
		sum -= nums[i]
	}
	return sum
}
