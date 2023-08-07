package main

import "fmt"

func sumOfUnique(nums []int) int {
	m := make([]int, 100)
	res := 0
	for _, val := range nums {
		m[val]++
	}

	for i, val := range m {
		if val == 1 {
			res += i
		}
	}

	return res
}

func main() {
	fmt.Println(sumOfUnique([]int{1, 2, 3, 2}))
}
