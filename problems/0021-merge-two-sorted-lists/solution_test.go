package tasks0021

import (
	. "leetcode/tools/structures/linked_list"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"test 01", args{ConvertArrayToList([]int{1, 2, 4}), ConvertArrayToList([]int{1, 3, 4})}, ConvertArrayToList([]int{1, 1, 2, 3, 4, 4})},
		{"test 02", args{ConvertArrayToList([]int{}), ConvertArrayToList([]int{})}, ConvertArrayToList([]int{})},
		{"test 03", args{ConvertArrayToList([]int{}), ConvertArrayToList([]int{0})}, ConvertArrayToList([]int{0})},
		{"test 04", args{ConvertArrayToList([]int{1}), ConvertArrayToList([]int{})}, ConvertArrayToList([]int{1})},
		{"test 05", args{ConvertArrayToList([]int{1}), ConvertArrayToList([]int{2})}, ConvertArrayToList([]int{1, 2})},
		{"test 06", args{ConvertArrayToList([]int{2}), ConvertArrayToList([]int{1})}, ConvertArrayToList([]int{1, 2})},
		{"test 07", args{ConvertArrayToList([]int{5}), ConvertArrayToList([]int{1, 2, 3})}, ConvertArrayToList([]int{1, 2, 3, 5})},
		{"test 08", args{ConvertArrayToList([]int{1, 2, 3}), ConvertArrayToList([]int{5})}, ConvertArrayToList([]int{1, 2, 3, 5})},
		{"test 09", args{ConvertArrayToList([]int{1, 2, 3, 4}), ConvertArrayToList([]int{5, 6, 7, 8})}, ConvertArrayToList([]int{1, 2, 3, 4, 5, 6, 7, 8})},
		{"test 10", args{ConvertArrayToList([]int{5, 6, 7, 8}), ConvertArrayToList([]int{1, 2, 3, 4})}, ConvertArrayToList([]int{1, 2, 3, 4, 5, 6, 7, 8})},
		{"test 11", args{ConvertArrayToList([]int{1, 1, 1}), ConvertArrayToList([]int{1, 1, 1})}, ConvertArrayToList([]int{1, 1, 1, 1, 1, 1})},
		{"test 12", args{ConvertArrayToList([]int{-5, 0, 5}), ConvertArrayToList([]int{-3, 3})}, ConvertArrayToList([]int{-5, -3, 0, 3, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
