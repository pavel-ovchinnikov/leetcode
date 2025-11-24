package removelinkedlistelements

import (
	. "github.com/pavel-ovchinnikov/leetcode/tools/list_node"
)

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	cur1 := dummy
	cur2 := head

	for cur2 != nil {
		if cur2.Val == val {
			cur2 = cur2.Next
			continue
		}

		cur1.Next = cur2
		cur1 = cur1.Next
		cur2 = cur1.Next
	}

	cur1.Next = nil

	return dummy.Next
}
