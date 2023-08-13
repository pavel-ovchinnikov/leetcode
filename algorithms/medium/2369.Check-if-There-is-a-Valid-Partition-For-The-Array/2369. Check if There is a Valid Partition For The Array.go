package main

import "fmt"

func validPartition(nums []int) bool {
	if len(nums) == 1 {
		return false
	}

	dp := []bool{true, false, nums[0] == nums[1]}
	for i := 2; i < len(nums); i++ {
		cur := false
		if isValid2(nums[i-1], nums[i]) && dp[1] {
			cur = true
		} else if isValid3(nums[i-2], nums[i-1], nums[i]) && dp[0] {
			cur = true
		} else if isConsecutive(nums[i-2], nums[i-1], nums[i]) && dp[0] {
			cur = true
		}

		dp[0], dp[1], dp[2] = dp[1], dp[2], cur
	}
	return dp[2]
}
func isValid2(a, b int) bool {
	return a == b
}

func isValid3(a, b, c int) bool {
	return (a == b) && (b == c)
}

func isConsecutive(a, b, c int) bool {
	return (c-b == 1) && (b-a == 1)
}

func main() {
	fmt.Println(validPartition([]int{4, 4, 4, 5, 6}))
}
