package main

import "fmt"

func majorityElement(nums []int) []int {
	n := len(nums)
	hash := make(map[int]int)
	for _, num := range nums {
		hash[num]++
	}

	res := make([]int, 0, n/3)
	for num, count := range hash {
		if count > n/3 {
			res = append(res, num)
		}
	}
	return res
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{3, 2, 3, 1, 3, 2, 5, 6}))
}
