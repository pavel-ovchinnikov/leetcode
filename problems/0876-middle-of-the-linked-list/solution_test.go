package middleofthelinkedlist

import (
	"testing"

	. "github.com/pavel-ovchinnikov/leetcode/tools/list_node"
)

func Test_middleNode(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *ListNode
		want *ListNode
	}{
		{
			name: "Example 1",
			head: ConvertArrayToList([]int{1, 2, 3, 4, 5}),
			want: ConvertArrayToList([]int{3, 4, 5}),
		},
		{
			name: "Example 2",
			head: ConvertArrayToList([]int{1, 2, 3, 4, 5, 6}),
			want: ConvertArrayToList([]int{4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := middleNode(tt.head)
			if !AreListsEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
