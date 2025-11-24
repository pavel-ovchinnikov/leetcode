package middleofthelinkedlist

import (
	. "github.com/pavel-ovchinnikov/leetcode/tools/list_node"
)

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
