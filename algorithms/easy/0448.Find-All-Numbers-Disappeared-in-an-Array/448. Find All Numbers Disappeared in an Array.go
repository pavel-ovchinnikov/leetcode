package main

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	d := len(nums)
	res := make([]int, d)

	for _, n := range nums {
		res[n-1] = n
	}
	j := 0
	for i := 0; i < d; i++ {
		if res[i] == 0 {
			res[j] = i + 1
			j++
		}
	}
	return res[:j]
}

func main() {
	fmt.Println([]int{4, 3, 2, 7, 8, 2, 3, 1})
}
