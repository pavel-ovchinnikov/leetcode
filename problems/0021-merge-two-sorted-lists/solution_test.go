package mergetwosortedlists

import (
	"testing"

	. "github.com/pavel-ovchinnikov/leetcode/tools/list_node"
)

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{
			name:  "Test Case 1",
			list1: ConvertArrayToList([]int{1, 2, 4}),
			list2: ConvertArrayToList([]int{1, 3, 4}),
			want:  ConvertArrayToList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name:  "Test Case 2",
			list1: nil,
			list2: nil,
			want:  nil,
		},
		{
			name:  "Test Case 3",
			list1: nil,
			list2: &ListNode{Val: 0},
			want:  &ListNode{Val: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.list1, tt.list2)
			if !AreListsEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
