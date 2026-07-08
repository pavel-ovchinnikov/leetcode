package tasks0083

import (
	. "leetcode/tools/structures/linked_list"
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"nil head", args{nil}, nil},
		{"single node", args{ConvertArrayToList([]int{1})}, ConvertArrayToList([]int{1})},
		{"no duplicates", args{ConvertArrayToList([]int{1, 2, 3})}, ConvertArrayToList([]int{1, 2, 3})},
		{"all duplicates", args{ConvertArrayToList([]int{1, 1, 1})}, ConvertArrayToList([]int{1})},
		{"duplicates at start", args{ConvertArrayToList([]int{1, 1, 2})}, ConvertArrayToList([]int{1, 2})},
		{"duplicates at end", args{ConvertArrayToList([]int{1, 2, 2})}, ConvertArrayToList([]int{1, 2})},
		{"duplicates in middle", args{ConvertArrayToList([]int{1, 2, 2, 3})}, ConvertArrayToList([]int{1, 2, 3})},
		{"multiple duplicate runs", args{ConvertArrayToList([]int{1, 1, 2, 3, 3})}, ConvertArrayToList([]int{1, 2, 3})},
		{"all same value many", args{ConvertArrayToList([]int{1, 1, 1, 1})}, ConvertArrayToList([]int{1})},
		{"negative numbers with duplicates", args{ConvertArrayToList([]int{-1, -1, 0, 1, 1})}, ConvertArrayToList([]int{-1, 0, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
