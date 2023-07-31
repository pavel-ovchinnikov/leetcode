package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	set := make(map[int]int)
	for i, num := range nums {
		v, ok := set[num]
		if ok && abs(i-v) <= k {
			return true
		}
		set[num] = i
	}
	return false
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
}
