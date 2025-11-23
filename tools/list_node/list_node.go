package listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

func ConvertArrayToList(values []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for _, val := range values {
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}

	return dummy.Next
}

func ConvertListToArray(head *ListNode) []int {
	result := make([]int, 0)
	cur := head
	for cur != nil {
		result = append(result, cur.Val)
		cur = cur.Next
	}

	return result
}

func AreListsEqual(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	return l1 == nil && l2 == nil
}
