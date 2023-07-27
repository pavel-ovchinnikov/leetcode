package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type ListNode = structures.ListNode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var cur *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = cur
		cur = head
		head = tmp
	}
	return cur
}

func main() {
	head := structures.Ints2List([]int{1, 2, 3, 4, 5})
	fmt.Println(structures.List2Ints(reverseList(head)))
}
