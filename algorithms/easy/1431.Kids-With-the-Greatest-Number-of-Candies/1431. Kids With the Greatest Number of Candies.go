package main

import "fmt"

func max(nums []int) (max int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandie := max(candies)

	result := make([]bool, len(candies))
	for i, candie := range candies {
		result[i] = maxCandie <= candie+extraCandies
	}
	return result
}

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
	fmt.Println(kidsWithCandies([]int{12, 1, 12}, 10))
}
