package addtwonumbers

import (
	"reflect"
	"testing"

	. "github.com/pavel-ovchinnikov/leetcode/tools/linked_list"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test 1",
			args{
				ConvertArrayToList([]int{2, 4, 3}),
				ConvertArrayToList([]int{5, 6, 4}),
			},
			ConvertArrayToList([]int{7, 0, 8}),
		},
		{
			"test 2",
			args{
				ConvertArrayToList([]int{9, 9, 9, 9, 9, 9, 9}),
				ConvertArrayToList([]int{9, 9, 9, 9}),
			},
			ConvertArrayToList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
