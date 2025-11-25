package swapnodesinpairs

import (
	. "github.com/pavel-ovchinnikov/leetcode/tools/list_node"
)

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{}
	cur := dummy
	l1 := head

	for l1 != nil && l1.Next != nil {
		l2 := l1.Next
		next := l2.Next

		cur.Next = l2
		l2.Next = l1
		l1.Next = next

		cur = l1
		l1 = next
	}

	return dummy.Next
}
