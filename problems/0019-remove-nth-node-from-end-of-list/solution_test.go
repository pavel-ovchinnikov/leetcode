package removenthnodefromendoflist

import (
	"testing"

	. "github.com/pavel-ovchinnikov/leetcode/tools/linked_list"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test 1", args{ConvertArrayToList([]int{1, 2, 3, 4, 5}), 2}, ConvertArrayToList([]int{1, 2, 3, 5})},
		{"test 2", args{ConvertArrayToList([]int{1}), 1}, ConvertArrayToList([]int{})},
		{"test 3", args{ConvertArrayToList([]int{1, 2}), 1}, ConvertArrayToList([]int{1})},
		{"test 4", args{ConvertArrayToList([]int{1, 2}), 2}, ConvertArrayToList([]int{2})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !Equal(got, tt.want) {
				t.Errorf(
					"removeNthFromEnd() = %v, want %v",
					ConvertListToArray(got),
					ConvertListToArray(tt.want),
				)
			}
		})
	}
}
