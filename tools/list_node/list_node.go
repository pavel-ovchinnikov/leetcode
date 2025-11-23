package listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

func convertArrayToList(values []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for _, val := range values {
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}

	return dummy.Next
}

func convertListToArray(head *ListNode) []int {
	result := make([]int, 0)
	cur := head
	for cur != nil {
		result = append(result, cur.Val)
		cur = cur.Next
	}

	return result
}
