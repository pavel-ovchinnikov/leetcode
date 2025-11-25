package swapnodesinpairs

import (
	"testing"

	. "github.com/pavel-ovchinnikov/leetcode/tools/list_node"
)

func Test_swapPairs(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		head *ListNode
		want *ListNode
	}{
		{
			name: "Test 1",
			head: ConvertArrayToList([]int{1, 2, 3, 4}),
			want: ConvertArrayToList([]int{2, 1, 4, 3}),
		},
		{
			name: "Test 2",
			head: ConvertArrayToList([]int{1, 2, 3}),
			want: ConvertArrayToList([]int{2, 1, 3}),
		},
		{
			name: "Test 3",
			head: ConvertArrayToList([]int{1}),
			want: ConvertArrayToList([]int{1}),
		},
		{
			name: "Test 4",
			head: ConvertArrayToList([]int{}),
			want: ConvertArrayToList([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := swapPairs(tt.head)
			if !AreListsEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
