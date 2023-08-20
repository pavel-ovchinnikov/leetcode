package main

import "fmt"

func distributeCandies(candyType []int) int {
	counter := make(map[int]struct{})
	for _, candy := range candyType {
		counter[candy] = struct{}{}
	}

	limit := len(candyType) / 2
	if len(counter) >= limit {
		return limit
	}
	return len(counter)
}
func main() {
	fmt.Println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
}
