package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type ListNode = structures.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cur1 := headA
	cur2 := headB

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

func end(head *ListNode) *ListNode {
	res := head
	for res.Next != nil {
		res = res.Next
	}

	return res
}

func firstNumber(head *ListNode, val int) *ListNode {
	node := head
	for node != nil && node.Val != val {
		node = node.Next
	}
	return node
}

func main() {
	tail := structures.Ints2List([]int{8, 4, 5})
	listA := structures.Ints2List([]int{4, 1})
	listB := structures.Ints2List([]int{5, 6, 1})

	end(listA).Next = tail
	end(listB).Next = tail

	fmt.Println(structures.List2Ints(getIntersectionNode(listA, listB)))
}
