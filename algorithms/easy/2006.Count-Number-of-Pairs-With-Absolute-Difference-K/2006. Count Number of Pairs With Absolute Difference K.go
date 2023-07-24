package main

import "fmt"

func countKDifference(nums []int, k int) int {
	res := 0
	mm := make(map[int]int)

	for _, num := range nums {
		mm[num] += 1
	}

	for key, val := range mm {
		if v, contains := mm[key-k]; contains {
			res += val * v
		}
	}

	return res
}

func main() {
	fmt.Println(countKDifference([]int{1, 2, 2, 1}, 1))
	fmt.Println(countKDifference([]int{1, 3}, 3))
	fmt.Println(countKDifference([]int{3, 2, 1, 5, 4}, 2))
	fmt.Println(countKDifference([]int{2, 1, 1, 9, 6, 2, 10, 4, 1, 6}, 6))

}
