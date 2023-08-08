package main

import (
	"fmt"
	"sort"
)

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	res := 0
	boxCount := 0
	for _, item := range boxTypes {
		if boxCount+item[0] >= truckSize {
			res += (truckSize - boxCount) * item[1]
			break
		} else {
			boxCount += item[0]
			res += item[0] * item[1]
		}
	}
	return res
}

func main() {
	fmt.Println(maximumUnits([][]int{
		{1, 3},
		{2, 2},
		{3, 1}},
		4))
}
