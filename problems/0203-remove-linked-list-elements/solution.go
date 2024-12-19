package mergetwosortedlists

import (
	. "github.com/pavel-ovchinnikov/leetcode/tools/structures/linked_list"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{}
	cur1, cur2 := newHead, head
	for cur2 != nil {
		if cur2.Val == val {
			cur2 = cur2.Next
			continue
		}

		cur1.Next = cur2
		cur1, cur2 = cur1.Next, cur2.Next
	}
	cur1.Next = nil
	return newHead.Next
}
