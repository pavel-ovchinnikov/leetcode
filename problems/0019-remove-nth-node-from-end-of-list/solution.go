package removenthnodefromendoflist

import (
	. "github.com/pavel-ovchinnikov/leetcode/tools/linked_list"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	p1, p2 := dummy, dummy

	for i := 0; i < n; i++ {
		p2 = p2.Next
	}

	for p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	if p1.Next != nil {
		p1.Next = p1.Next.Next
	}

	return dummy.Next
}
