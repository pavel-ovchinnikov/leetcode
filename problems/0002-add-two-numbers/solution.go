package addtwonumbers

import (
	. "github.com/pavel-ovchinnikov/leetcode/tools/linked_list"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	cur := root
	d := 0

	for l1 != nil && l2 != nil {
		d += l1.Val + l2.Val
		cur.Next = &ListNode{Val: d % 10}
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
		d /= 10
	}

	for l1 != nil {
		d += l1.Val
		cur.Next = &ListNode{Val: d % 10}
		cur = cur.Next
		l1 = l1.Next
		d /= 10
	}

	for l2 != nil {
		d += l2.Val
		cur.Next = &ListNode{Val: d % 10}
		cur = cur.Next
		l2 = l2.Next
		d /= 10
	}

	if d != 0 {
		cur.Next = &ListNode{Val: d}
	}

	return root.Next
}
