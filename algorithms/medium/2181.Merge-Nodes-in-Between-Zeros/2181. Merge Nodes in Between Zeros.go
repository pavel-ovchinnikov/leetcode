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

func mergeNodes(head *ListNode) *ListNode {
	root := &ListNode{}
	cur2 := root

	cur := head
	sum := 0

	for cur != nil {
		if cur.Val == 0 && sum != 0 {
			cur2.Next = &ListNode{Val: sum}
			cur2 = cur2.Next
			sum = 0
		} else {
			sum += cur.Val
		}
		cur = cur.Next
	}

	return root.Next
}

func main() {
	head := structures.Ints2List([]int{0, 3, 1, 0, 4, 5, 2, 0})
	fmt.Println(structures.List2Ints(mergeNodes(head)))
}
