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

func partition(head *ListNode, x int) *ListNode {
	head1, head2 := &ListNode{}, &ListNode{}
	cur1, cur2 := head1, head2
	cur := head

	for cur != nil {
		fmt.Println(cur)
		if cur.Val < x {
			cur1.Next = cur
			cur1 = cur1.Next

		} else {
			cur2.Next = cur
			cur2 = cur2.Next
		}
		cur = cur.Next
	}

	cur1.Next = head2.Next
	cur2.Next = nil
	return head1.Next
}

func main() {
	head := structures.Ints2List([]int{1, 4, 3, 2, 5, 2})
	fmt.Println(structures.List2Ints(partition(head, 3)))

}
