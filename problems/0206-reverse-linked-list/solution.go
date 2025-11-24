package reverselinkedlist

import (
	. "github.com/pavel-ovchinnikov/leetcode/tools/list_node"
)

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev

		prev = curr
		curr = next
	}

	return prev
}
