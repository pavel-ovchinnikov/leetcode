package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	hash := map[int]int{}
	for _, v := range arr {
		hash[v]++
	}
	hash2 := map[int]bool{}
	for _, v := range hash {
		hash2[v] = true
	}

	return len(hash) == len(hash2)
}

func main() {
	fmt.Println(uniqueOccurrences([]int{1, 1, 2, 2, 2, 3, 3, 3, 3}))
}
