package main

import "fmt"

func maxSum(nums []int) int {
	hash := map[int]int{}
	maxSum := -1
	for _, num := range nums {
		maxDigit := MaxDigit(num)
		if val, ok := hash[maxDigit]; ok {
			maxSum = max(maxSum, val+num)
			hash[maxDigit] = max(num, val)
		} else {
			hash[maxDigit] = num
		}
	}
	return maxSum
}

func MaxDigit(num int) int {
	maxDigit := num % 10
	for num != 0 {
		maxDigit = max(maxDigit, num%10)
		num /= 10
	}
	return maxDigit
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(maxSum([]int{51, 71, 17, 24, 42}))
	fmt.Println(maxSum([]int{1, 2, 3, 4}))

}
