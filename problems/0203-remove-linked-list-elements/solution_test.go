package removelinkedlistelements

import (
	"testing"

	. "github.com/pavel-ovchinnikov/leetcode/tools/list_node"
)

func Test_removeElements(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *ListNode
		val  int
		want *ListNode
	}{
		{
			name: "",
			head: ConvertArrayToList([]int{1, 2, 6, 3, 4, 5, 6}),
			val:  6,
			want: ConvertArrayToList([]int{1, 2, 3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElements(tt.head, tt.val)
			if !AreListsEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
