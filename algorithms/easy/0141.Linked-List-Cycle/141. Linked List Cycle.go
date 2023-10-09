package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type ListNode = structures.ListNode

// func hasCycle(head *ListNode) bool {
// 	if head == nil {
// 		return false
// 	}
// 	cur1 := head
// 	cur2 := head.Next

// 	for cur1 != nil && cur2 != nil && cur2.Next != nil {
// 		if cur1 == cur2 {
// 			return true
// 		}
// 		cur1 = cur1.Next
// 		cur2 = cur2.Next.Next
// 	}
// 	return false
// }

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	l1, l2 := head, head.Next

	for l2 != nil && l2.Next != nil && l1.Next.Next != nil && l1 != l2 {
		l1 = l1.Next
		l2 = l2.Next.Next
	}

	return l1 == l2
}

func main() {
	head := structures.Ints2ListWithCycle([]int{3, 2, 0, -4}, 1)
	fmt.Println(hasCycle(head))
}
