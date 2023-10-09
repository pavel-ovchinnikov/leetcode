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

func mergeKLists(lists []*ListNode) *ListNode {
	var res *ListNode
	for i := 0; i < len(lists); i++ {
		res = mergeLists(res, lists[i])
	}
	return res
}

func mergeLists(a, b *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for a != nil && b != nil {
		if a.Val < b.Val {
			cur.Next = a
			a = a.Next
		} else {
			cur.Next = b
			b = b.Next
		}
		cur = cur.Next
	}

	if a != nil {
		cur.Next = a
	}

	if b != nil {
		cur.Next = b
	}

	return head.Next
}

func main() {
	lists := []*ListNode{
		structures.Ints2List([]int{1, 4, 5}),
		structures.Ints2List([]int{1, 4, 5}),
		structures.Ints2List([]int{2, 6}),
	}
	fmt.Println(structures.List2Ints(mergeKLists(lists)))
}
