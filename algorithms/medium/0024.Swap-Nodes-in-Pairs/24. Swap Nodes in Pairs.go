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

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	root := head.Next
	head.Next = swapPairs(head.Next.Next)
	root.Next = head

	return root
}

func main() {
	head := structures.Ints2List([]int{1, 2, 3, 4})
	fmt.Println(structures.List2Ints(swapPairs(head)))

}
