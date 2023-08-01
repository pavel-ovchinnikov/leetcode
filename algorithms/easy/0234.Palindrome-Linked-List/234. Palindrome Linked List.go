package main

import (
	"fmt"

	"github.com/pavel-ovchinnikov/LeetCode/structures"
)

type ListNode = structures.ListNode

func isPalindrome(head *ListNode) bool {
	it1 := head
	it2 := reverseList(endOfFirstHalf(head).Next)

	for it1 != nil && it2 != nil && it1.Val == it2.Val {
		it1 = it1.Next
		it2 = it2.Next
	}
	return it2 == nil
}

func endOfFirstHalf(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func main() {
	head := structures.Ints2List([]int{1, 2, 2, 1})
	fmt.Println(isPalindrome(head))
}
