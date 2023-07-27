package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

// ListNode define
type ListNode = structures.ListNode

func getDecimalValue(head *ListNode) int {
	result := 0
	for head != nil {
		result = (result << 1) + head.Val
		head = head.Next
	}
	return result
}

func main() {
	head := structures.Ints2List([]int{1, 0, 1})
	fmt.Println(getDecimalValue(head))
}
