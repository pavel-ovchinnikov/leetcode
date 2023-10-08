package main

import "fmt"

func numEquivDominoPairs(dominoes [][]int) int {
	hash := make([][]int, 10)
	res := 0

	for i := 0; i < 10; i++ {
		hash[i] = make([]int, 10)
	}

	for i := 0; i < len(dominoes); i++ {
		begin, end := dominoes[i][0], dominoes[i][1]

		res += hash[begin][end]

		if begin != end {
			res += hash[end][begin]
		}

		hash[begin][end]++
	}

	return res
}

func main() {
	fmt.Println(numEquivDominoPairs([][]int{
		{1, 2},
		{1, 2},
		{1, 1},
		{1, 2},
		{2, 2},
	}))
}
