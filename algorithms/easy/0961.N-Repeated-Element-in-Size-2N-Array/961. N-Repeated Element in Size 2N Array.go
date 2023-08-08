package main

import "fmt"

func repeatedNTimes(nums []int) int {
	hash := make(map[int]bool)

	for _, v := range nums {
		_, ok := hash[v]
		if ok {
			return v
		}

		hash[v] = true
	}

	return -1
}

func main() {
	fmt.Println(repeatedNTimes([]int{1, 2, 3, 3}))
}
