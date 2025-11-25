package intersectionoftwolinkedlists

import (
	"testing"

	. "github.com/pavel-ovchinnikov/leetcode/tools/list_node"
)

func Test_getIntersectionNode(t *testing.T) {
	tail := ConvertArrayToList([]int{})
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		headA *ListNode
		headB *ListNode
		want  *ListNode
	}{
		{
			name: "",
			headA: &ListNode{
				Val:  10,
				Next: tail,
			},
			headB: &ListNode{
				Val:  100,
				Next: tail,
			},
			want: tail,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getIntersectionNode(tt.headA, tt.headB)
			if !AreListsEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
