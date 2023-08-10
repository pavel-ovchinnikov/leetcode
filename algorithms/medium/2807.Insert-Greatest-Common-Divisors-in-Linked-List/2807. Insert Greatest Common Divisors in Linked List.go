package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type ListNode = structures.ListNode

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		node := &ListNode{Val: gcd(cur.Val, cur.Next.Val)}
		cur.Next, node.Next = node, cur.Next
		cur = node.Next
	}
	return head
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func main() {
	{
		head := structures.Ints2List([]int{18, 6, 10, 3})
		out := structures.List2Ints(insertGreatestCommonDivisors(head))
		fmt.Println(out)
	}
	{
		head := structures.Ints2List([]int{18, 6, 3})
		out := structures.List2Ints(insertGreatestCommonDivisors(head))
		fmt.Println(out)
	}
	{
		head := structures.Ints2List([]int{18})
		out := structures.List2Ints(insertGreatestCommonDivisors(head))
		fmt.Println(out)
	}
	{
		head := structures.Ints2List([]int{})
		out := structures.List2Ints(insertGreatestCommonDivisors(head))
		fmt.Println(out)
	}

}
