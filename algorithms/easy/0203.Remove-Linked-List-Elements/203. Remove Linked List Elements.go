package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type ListNode = structures.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	root := &ListNode{}
	cur := root

	for head != nil {
		if head.Val != val {
			cur.Next = head
			cur = cur.Next
		}
		head = head.Next
	}

	cur.Next = nil

	return root.Next
}

func main() {
	{
		head := structures.Ints2List([]int{1, 2, 6, 3, 4, 5, 6})
		fmt.Println(structures.List2Ints(removeElements(head, 6)))
	}
	{
		head := structures.Ints2List([]int{1, 2, 6, 3, 4, 5, 6})
		fmt.Println(structures.List2Ints(removeElements(head, 1)))
	}
	{
		head := structures.Ints2List([]int{1})
		fmt.Println(structures.List2Ints(removeElements(head, 1)))
	}

}
