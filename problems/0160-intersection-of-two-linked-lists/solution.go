package intersectionoftwolinkedlists

import (
	. "github.com/pavel-ovchinnikov/leetcode/tools/structures/linked_list"
)

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	diff := length(headA) - length(headB)
	cur1, cur2 := headA, headB
	if diff < 0 {
		cur1, cur2 = headB, headA
		diff = -diff
	}

	for range diff {
		cur1 = cur1.Next
	}

	for cur1 != cur2 {
		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	return cur1
}

func length(head *ListNode) int {
	l := 0
	for head != nil {
		head = head.Next
		l++
	}
	return l
}

func getIntersectionNode_v2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	cur1, cur2 := headA, headB
	for cur1 != cur2 {
		if cur1 == nil {
			cur1 = headB
		} else {
			cur1 = cur1.Next
		}

		if cur2 == nil {
			cur2 = headA
		} else {
			cur2 = cur2.Next
		}
	}

	return cur1
}
