package tasks0083

import (
	. "leetcode/tools/structures/linked_list"
)

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	curr := head
	for curr != nil && prev != nil {
		if prev.Val == curr.Val {
			prev.Next = curr.Next
			temp := curr.Next
			curr.Next = nil
			curr = temp
		} else {
			curr = curr.Next
			prev = prev.Next
		}
	}
	return dummy.Next
}
