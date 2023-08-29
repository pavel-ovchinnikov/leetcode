package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = structures.ListNode

func pairSum(head *ListNode) int {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	var nextNode *ListNode
	var prev *ListNode

	for slow != nil {
		nextNode = slow.Next
		slow.Next = prev
		prev, slow = slow, nextNode
	}

	res := 0
	for prev != nil {
		res = max(res, head.Val+prev.Val)
		prev = prev.Next
		head = head.Next
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	head := structures.Ints2List([]int{4, 2, 2, 3})
	fmt.Println(pairSum(head))
}
