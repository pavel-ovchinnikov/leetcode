package convertbinarynumberinalinkedlisttointeger

import (
	. "github.com/pavel-ovchinnikov/leetcode/tools/list_node"
)

func getDecimalValue(head *ListNode) int {
	cur := head
	result := 0

	for cur != nil {
		result = result<<1 + cur.Val
		cur = cur.Next
	}

	return result
}
