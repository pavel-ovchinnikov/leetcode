package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type ListNode = structures.ListNode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	root := &ListNode{}
	curr := root

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 == nil {
		curr.Next = list2
	} else {
		curr.Next = list1
	}

	return root.Next
}

func main() {
	list1 := structures.Ints2List([]int{1, 2, 4})
	list2 := structures.Ints2List([]int{1, 3, 4})
	fmt.Println(structures.List2Ints(mergeTwoLists(list1, list2)))
}
