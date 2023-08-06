package main

import "fmt"

func distinctDifferenceArray(nums []int) []int {
	res := make([]int, len(nums))
	prefix := make(map[int]int)
	suffix := make(map[int]int)

	for _, num := range nums {
		suffix[num] += 1
	}

	for i, num := range nums {
		prefix[num]++
		suffix[num] -= 1
		if suffix[num] == 0 {
			delete(suffix, num)
		}

		res[i] = len(prefix) - len(suffix)
	}

	return res
}

func main() {
	fmt.Println(distinctDifferenceArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(distinctDifferenceArray([]int{1, 1, 1, 1, 2}))
	fmt.Println(distinctDifferenceArray([]int{1, 1, 1, 1, 3}))
}
