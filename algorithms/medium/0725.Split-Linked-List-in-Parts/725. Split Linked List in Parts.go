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

func splitListToParts(head *ListNode, k int) []*ListNode {
	size := count(head)
	n := size / k
	d := size % k
	cur := head
	res := make([]*ListNode, k)

	for i := 0; i < k && cur != nil; i++ {
		res[i] = cur
		prev := cur

		for j := 0; j < n && cur != nil; j++ {
			prev = cur
			cur = cur.Next
		}

		if d > 0 && cur != nil {
			prev = cur
			cur = cur.Next
			d--
		}

		prev.Next = nil
	}

	return res
}

func count(head *ListNode) (res int) {
	for head != nil {
		res++
		head = head.Next
	}
	return
}

func main() {
	// head := structures.Ints2List([]int{1, 2, 3})
	head := structures.Ints2List([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	res := splitListToParts(head, 3)
	for _, node := range res {
		if node == nil {
			fmt.Println("[]")
		} else {
			for node != nil {
				fmt.Print(node.Val, " ")
				node = node.Next
			}
			fmt.Println()
		}

	}
}
