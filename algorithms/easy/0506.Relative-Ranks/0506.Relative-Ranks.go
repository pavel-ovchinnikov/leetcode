package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Pair struct {
	Val   int
	Index int
}

func findRelativeRanks(nums []int) []string {
	pairs := make([]Pair, len(nums))
	for i, num := range nums {
		pairs[i].Index = i
		pairs[i].Val = num
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Val > pairs[j].Val
	})
	res := make([]string, len(nums))

	for i, athlete := range pairs {
		switch i {
		case 0:
			res[athlete.Index] = "Gold Medal"
		case 1:
			res[athlete.Index] = "Silver Medal"
		case 2:
			res[athlete.Index] = "Bronze Medal"
		default:
			res[athlete.Index] = strconv.Itoa(i + 1)
		}
	}

	return res
}

func main() {
	fmt.Println(findRelativeRanks([]int{5, 4, 3, 2, 1}))
}
