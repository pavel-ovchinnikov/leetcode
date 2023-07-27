package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type ListNode = structures.ListNode

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	mid := head
	cur := head
	for cur != nil {

		if cur.Next == nil {
			return mid
		}
		cur = cur.Next
		cur = cur.Next
		mid = mid.Next
	}
	return mid
}

func main() {
	{
		head := structures.Ints2List([]int{1, 2, 3, 4, 5})
		fmt.Println(structures.List2Ints(middleNode(head)))
	}
	{
		head := structures.Ints2List([]int{1, 2, 3, 4, 5})
		fmt.Println(structures.List2Ints(middleNode(head)))
	}

}
