package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = structures.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	root := &ListNode{}
	root.Next = head
	prev := root
	cur := head

	for i := 0; i < n; i++ {
		cur = cur.Next
	}
	for cur != nil {
		cur = cur.Next
		prev = prev.Next
	}

	prev.Next = prev.Next.Next

	return root.Next
}

func main() {
	head := structures.Ints2List([]int{1, 2, 3, 4, 5})
	fmt.Println(structures.List2Ints(removeNthFromEnd(head, 2)))
}
