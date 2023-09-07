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

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	if left == right {
		return head
	}

	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	cur := dummy
	prev := cur

	for i := 0; i < left-1; i++ {
		cur = cur.Next
	}

	prev = cur
	cur = cur.Next

	for i := 0; i < right-left; i++ {
		tmp := prev.Next
		prev.Next = cur.Next
		cur.Next = cur.Next.Next
		prev.Next.Next = tmp
	}

	return dummy.Next
}

func main() {
	head := structures.Ints2List([]int{1, 2, 3, 4, 5})
	fmt.Println(structures.List2Ints(reverseBetween(head, 2, 4)))
}
