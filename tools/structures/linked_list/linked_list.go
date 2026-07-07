package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func ConvertArrayToList(nums []int) *ListNode {
	root := &ListNode{}
	cur := root
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}

	return root.Next
}

func ConvertArrayToListWithCycle(nums []int, pos int) *ListNode {
	head := ConvertArrayToList(nums)
	end := endList(head)
	node := getNodeByPos(head, pos)
	end.Next = node
	return head
}

func ConvertListToArray(root *ListNode) []int {
	arr := make([]int, 0)
	for root != nil {
		arr = append(arr, root.Val)
		root = root.Next
	}
	return arr
}

func endList(head *ListNode) *ListNode {
	if head == nil {
		panic("head is nil")
	}

	for head != nil && head.Next != nil {
		head = head.Next
	}

	return head
}

func getNodeByPos(head *ListNode, pos int) *ListNode {
	if head == nil {
		panic("head is nil")
	}

	cur := head
	for i := 0; i < pos; i++ {
		cur = cur.Next
	}

	return cur
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
