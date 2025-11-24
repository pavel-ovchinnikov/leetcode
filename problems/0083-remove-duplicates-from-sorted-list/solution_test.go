package removeduplicatesfromsortedlist

import (
	"reflect"
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
			name: "Test Case 1",
			head: ConvertArrayToList([]int{1, 1, 2}),
			want: ConvertArrayToList([]int{1, 2}),
		},
		{
			name: "Test Case 2",
			head: ConvertArrayToList([]int{1, 1, 2, 3, 3}),
			want: ConvertArrayToList([]int{1, 2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := deleteDuplicates(tt.head)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
