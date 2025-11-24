package removeduplicatesfromsortedlist

import (
	. "github.com/pavel-ovchinnikov/leetcode/tools/list_node"
)

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Val: head.Val}
	cur := dummy
	node := head.Next

	for node != nil {
		if cur.Val == node.Val {
			node = node.Next
			continue
		}
		cur.Next = node
		cur = cur.Next
		node = node.Next
	}

	cur.Next = nil

	return dummy
}
