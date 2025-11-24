package reverselinkedlist

import (
	"testing"

	. "github.com/pavel-ovchinnikov/leetcode/tools/list_node"
)

func Test_reverseList(t *testing.T) {

	tests := []struct {
		name string // description of this test case
		head *ListNode
		want *ListNode
	}{
		{
			name: "Task 1",
			head: ConvertArrayToList([]int{1, 2, 3, 4, 5}),
			want: ConvertArrayToList([]int{5, 4, 3, 2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList(tt.head)
			if !AreListsEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
