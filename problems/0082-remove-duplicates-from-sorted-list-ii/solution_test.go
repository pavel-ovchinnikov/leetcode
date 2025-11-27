package removeduplicatesfromsortedlistii

import (
	"testing"

	. "github.com/pavel-ovchinnikov/leetcode/tools/list_node"
)

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *ListNode
		want *ListNode
	}{
		{
			name: "",
			head: ConvertArrayToList([]int{1, 2, 3, 3, 4, 4, 5}),
			want: ConvertArrayToList([]int{1, 2, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteDuplicates(tt.head)
			if !AreListsEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", ConvertListToArray(got), ConvertListToArray(tt.want))
			}
		})
	}
}
