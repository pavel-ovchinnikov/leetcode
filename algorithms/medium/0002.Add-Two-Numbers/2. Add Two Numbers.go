package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type ListNode = structures.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	cur := root
	tmp := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			tmp += l1.Val
		}
		if l2 != nil {
			tmp += l2.Val
		}
		cur.Next = &ListNode{Val: tmp % 10, Next: nil}
		cur = cur.Next
		tmp /= 10

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if tmp != 0 {
		cur.Next = &ListNode{Val: tmp % 10, Next: nil}
	}
	return root.Next
}

func main() {
	num1 := structures.Ints2List([]int{2, 4, 3})
	num2 := structures.Ints2List([]int{5, 6, 4})
	fmt.Println(structures.List2Ints(addTwoNumbers(num1, num2)))
}
