package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func ConvertArrayToList(nums []int) *ListNode {
	head := &ListNode{}
	cur := head
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}

	return head.Next
}

func ConvertListToArray(node *ListNode) []int {
	arr := make([]int, 0)
	for node != nil {
		arr = append(arr, node.Val)
		node = node.Next
	}
	return arr
}

func Equal(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a, b = a.Next, b.Next
	}

	return a == nil && b == nil
}
