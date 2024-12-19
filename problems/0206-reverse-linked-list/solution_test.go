package reverselinkedlist

import (
	"reflect"
	"testing"

	. "github.com/pavel-ovchinnikov/leetcode/tools/structures/linked_list"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test 1", args{ConvertArrayToList([]int{1, 2, 3, 4, 5})}, ConvertArrayToList([]int{5, 4, 3, 2, 1})},
		{"test 2", args{ConvertArrayToList([]int{1, 2})}, ConvertArrayToList([]int{2, 1})},
		{"test 3", args{ConvertArrayToList([]int{1})}, ConvertArrayToList([]int{1})},
		{"test 4", args{ConvertArrayToList([]int{2, 1, 1})}, ConvertArrayToList([]int{1, 1, 2})},
		{"test 5", args{ConvertArrayToList([]int{1, 1, 1})}, ConvertArrayToList([]int{1, 1, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
