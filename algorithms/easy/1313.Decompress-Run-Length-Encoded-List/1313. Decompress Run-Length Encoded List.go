package main

import "fmt"

func decompressRLElist(nums []int) []int {
	var res []int

	for i := 0; i < len(nums); i += 2 {
		for count := 0; count < nums[i]; count++ {
			res = append(res, nums[i+1])
		}
	}

	return res
}

func main() {
	fmt.Println(decompressRLElist([]int{1, 2, 3, 4}))
	fmt.Println(decompressRLElist([]int{1, 1, 2, 3}))
}
