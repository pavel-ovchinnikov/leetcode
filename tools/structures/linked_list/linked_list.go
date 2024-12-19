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
