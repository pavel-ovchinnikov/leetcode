package main

import (
	"fmt"
	"sort"
)

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] == pairs[j][0] {
			return pairs[i][1] < pairs[j][1]
		}
		return pairs[i][0] < pairs[j][0]
	})

	res := 1
	j := 0
	for i := 1; i < len(pairs); i++ {
		if pairs[j][1] < pairs[i][0] {
			res++
			j = i
		} else {
			if pairs[j][1] > pairs[i][1] {
				j = i
			}
		}
	}

	return res
}

func main() {
	fmt.Println(findLongestChain([][]int{
		{1, 2},
		{7, 8},
		{4, 5},
	}))
}
