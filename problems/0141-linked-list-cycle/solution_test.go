package linkedlistcycle

import (
	"testing"

	. "github.com/pavel-ovchinnikov/leetcode/tools/list_node"
)

func Test_hasCycle(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *ListNode
		want bool
	}{
		{
			name: "Test case 1",
			head: func() *ListNode {
				node1 := &ListNode{Val: 1}
				node2 := &ListNode{Val: 2}
				node1.Next = node2
				node1.Next = node1 // Create a cycle here
				return node1
			}(),
			want: true,
		},
		{
			name: "Test case 2",
			head: func() *ListNode {
				node1 := &ListNode{Val: 3}
				node2 := &ListNode{Val: 2}
				node3 := &ListNode{Val: 0}
				node4 := &ListNode{Val: -4}
				node1.Next = node2
				node2.Next = node3
				node3.Next = node4
				node4.Next = node2 // Create a cycle here
				return node1
			}(),
			want: true,
		},
		{
			name: "Test case 3",
			head: ConvertArrayToList([]int{1, 2, 3, 4}),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasCycle(tt.head)
			if got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
