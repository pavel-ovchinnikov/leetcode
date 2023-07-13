package main

import "fmt"

func removeElement(nums []int, val int) int {
	index := 0
	for _, v := range nums {
		if v != val {
			nums[index] = v
			index++
		}
	}
	return index
}

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
}
