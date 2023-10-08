package main

import "fmt"

func canThreePartsEqualSum(arr []int) bool {
	sum := 0
	for _, num := range arr {
		sum += num
	}

	if sum%3 != 0 {
		return false
	}

	curSum, count := 0, 0
	for _, num := range arr {
		curSum += num

		if curSum == sum/3 {
			curSum = 0
			count++
		}
	}
	return count >= 3
}

func main() {
	fmt.Println(canThreePartsEqualSum([]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}))
}
