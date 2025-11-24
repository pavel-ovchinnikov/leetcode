package linkedlistcycle

import (
	. "github.com/pavel-ovchinnikov/leetcode/tools/list_node"
)

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	cur1, cur2 := head, head.Next
	for cur1 != nil && cur2.Next != nil && cur2.Next.Next != nil && cur1 != cur2 {
		cur1 = cur1.Next
		cur2 = cur2.Next.Next
	}

	return cur1 != nil && cur2 != nil && cur1 == cur2
}
