package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumIndicesWithKSetBits(nums []int, k int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		if hasBits(i, k) {
			sum += nums[i]
		}
	}

	return sum
}

func hasBits(n int, k int) bool {
	num := int64(n)
	nrconvstr := strconv.FormatInt(num, 2)
	return strings.Count(nrconvstr, "1") == k
}

func main() {
	fmt.Println(sumIndicesWithKSetBits([]int{5, 10, 1, 5, 2}, 2))
}
