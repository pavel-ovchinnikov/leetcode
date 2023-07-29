package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type ListNode = structures.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	root := head
	cur := root
	head = head.Next

	for head != nil {
		if cur.Val != head.Val {
			cur.Next = head
			cur = cur.Next
		}
		head = head.Next
	}

	cur.Next = nil

	return root
}

func main() {
	head := structures.Ints2List([]int{1, 1, 2})
	fmt.Println(structures.List2Ints(deleteDuplicates(head)))
}
