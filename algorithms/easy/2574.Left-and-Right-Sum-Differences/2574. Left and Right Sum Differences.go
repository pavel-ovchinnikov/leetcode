package main

import "fmt"

func leftRigthDifference(nums []int) []int {
	leftSum := 0
	rightSum := 0

	for _, num := range nums {
		rightSum += num
	}

	answer := make([]int, len(nums))
	for i := 0; i < len(answer); i++ {
		rightSum -= nums[i]
		answer[i] = absDiff(rightSum, leftSum)
		leftSum += nums[i]
	}

	return answer
}

func absDiff(num1, num2 int) int {
	diff := num1 - num2

	if diff < 0 {
		return -diff
	}

	return diff
}

func main() {
	fmt.Println(leftRigthDifference([]int{10, 4, 8, 3}))
}
