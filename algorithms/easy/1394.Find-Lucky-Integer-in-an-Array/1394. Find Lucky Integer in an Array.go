package main

import "fmt"

func findLucky(arr []int) int {
	hash := make(map[int]int)

	for _, num := range arr {
		hash[num]++
	}

	maxCount := -1
	for key, val := range hash {
		if key == val && maxCount < val {
			maxCount = val
		}
	}

	return maxCount
}

func main() {
	fmt.Println(findLucky([]int{1, 2, 2, 3, 3, 3}))
}
