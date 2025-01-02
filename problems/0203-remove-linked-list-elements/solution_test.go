package removelinkedlistelements

import (
	"testing"

	. "github.com/pavel-ovchinnikov/leetcode/tools/structures/linked_list"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test 1", args{ConvertArrayToList([]int{1, 1, 1}), 0}, []int{1, 1, 1}},
		{"test 2", args{ConvertArrayToList([]int{1, 2, 3}), 2}, []int{1, 3}},
		{"test 3", args{ConvertArrayToList([]int{1, 2, 3}), 1}, []int{2, 3}},
		{"test 4", args{ConvertArrayToList([]int{1, 1, 3}), 1}, []int{3}},
		{"test 5", args{ConvertArrayToList([]int{1, 3, 3}), 3}, []int{1}},
		{"test 6", args{ConvertArrayToList([]int{1, 2, 2, 2, 3}), 2}, []int{1, 3}},
		{"test 7", args{ConvertArrayToList([]int{}), 1}, []int{}},
		{"test 8", args{ConvertArrayToList([]int{7, 7, 7, 7}), 7}, []int{}},
		{"test 9", args{ConvertArrayToList([]int{1, 2, 6, 3, 4, 5, 6}), 6}, []int{1, 2, 3, 4, 5}},
		{"test 10", args{ConvertArrayToList([]int{1, 2, 3, 4, 1, 1, 2, 1, 1, 1, 6, 1, 1}), 1}, []int{2, 3, 4, 2, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !Equal(got, ConvertArrayToList(tt.want)) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
