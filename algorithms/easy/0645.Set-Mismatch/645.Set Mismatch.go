package main

import "fmt"

func findErrorNums(nums []int) []int {
	var sum, dup int
	n := len(nums)
	foundDuplicate := false
	for _, x := range nums {
		x1 := abs(x)
		sum += x1
		if !foundDuplicate {
			if nums[x1-1] < 0 {
				dup = x1
				foundDuplicate = true
			} else {
				nums[x1-1] = -nums[x1-1]
			}
		}
	}
	sumExpected, sumUniq := n*(n+1)/2, sum-dup
	return []int{dup, sumExpected - sumUniq}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
}
