package removeduplicatesfromsortedlist

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
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	cur1, prev := head.Next, head
	for cur1 != nil {
		if cur1.Val == prev.Val {
			cur1 = cur1.Next
			continue
		}
		prev.Next = cur1
		prev = prev.Next
		cur1 = cur1.Next
	}
	prev.Next = nil
	return head
}
