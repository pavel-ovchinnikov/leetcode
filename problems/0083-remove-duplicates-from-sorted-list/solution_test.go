package removeduplicatesfromsortedlist

import (
	"reflect"
	"testing"

	. "github.com/pavel-ovchinnikov/leetcode/tools/structures/linked_list"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test 1", args{ConvertArrayToList([]int{1, 1, 1})}, []int{1}},
		{"test 2", args{ConvertArrayToList([]int{1})}, []int{1}},
		{"test 3", args{ConvertArrayToList([]int{1, 1, 2})}, []int{1, 2}},
		{"test 4", args{ConvertArrayToList([]int{1, 2, 2, 2, 3})}, []int{1, 2, 3}},
		{"test 5", args{ConvertArrayToList([]int{1, 2, 3, 3})}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(ConvertListToArray(got), tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
