package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	a := make([]int, len(heights))
	res := 0
	copy(a, heights)
	sort.Ints(a)

	for i := 0; i < len(heights); i++ {
		if heights[i] != a[i] {
			res++
		}
	}
	return res

}

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3}))
}
