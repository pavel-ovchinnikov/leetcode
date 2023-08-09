package main

import (
	"fmt"
)

func countBeautifulPairs(nums []int) int {
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if GCD(firsDigit(nums[i]), nums[j]%10) == 1 {
				res++
			}
		}
	}
	return res
}

func firsDigit(nums int) int {
	for nums >= 10 {
		nums /= 10
	}
	return nums
}

func GCD(a, b int) int {
	for b != 0 {
		b, a = a%b, b
	}
	return a
}

func main() {
	fmt.Println(countBeautifulPairs([]int{2, 5, 1, 4}))
}
