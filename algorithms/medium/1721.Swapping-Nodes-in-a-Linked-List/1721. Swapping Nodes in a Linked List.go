package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type ListNode = structures.ListNode

func swapNodes(head *ListNode, k int) *ListNode {
	left, right := head, head
	cur := head
	for cur != nil {
		k--
		if k == 0 {
			left = cur
		} else if k < 0 {
			right = right.Next
		}

		cur = cur.Next
	}
	left.Val, right.Val = right.Val, left.Val
	return head
}

func main() {
	head := structures.Ints2List([]int{7, 9, 6, 6, 7, 8, 3, 0, 9, 5})
	fmt.Println(structures.List2Ints(swapNodes(head, 5)))
}
