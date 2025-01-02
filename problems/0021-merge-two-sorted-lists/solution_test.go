package mergetwosortedlists

import (
	"reflect"
	"testing"

	. "github.com/pavel-ovchinnikov/leetcode/tools/structures/linked_list"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test 1",
			args{
				ConvertArrayToList([]int{1, 2, 3}),
				ConvertArrayToList([]int{}),
			},
			ConvertArrayToList([]int{1, 2, 3}),
		},
		{
			"test 2",
			args{
				ConvertArrayToList([]int{}),
				ConvertArrayToList([]int{1, 2, 3}),
			},
			ConvertArrayToList([]int{1, 2, 3}),
		},
		{
			"test 3",
			args{
				ConvertArrayToList([]int{1, 2}),
				ConvertArrayToList([]int{1, 2, 3}),
			},
			ConvertArrayToList([]int{1, 1, 2, 2, 3}),
		},
		{
			"test 4",
			args{
				ConvertArrayToList([]int{3, 4}),
				ConvertArrayToList([]int{1, 2, 3}),
			},
			ConvertArrayToList([]int{1, 2, 3, 3, 4}),
		},
		{
			"test 5",
			args{
				ConvertArrayToList([]int{4}),
				ConvertArrayToList([]int{1, 2, 3}),
			},
			ConvertArrayToList([]int{1, 2, 3, 4}),
		},
		{
			"test 6",
			args{
				ConvertArrayToList([]int{1, 2, 3}),
				ConvertArrayToList([]int{1, 2}),
			},
			ConvertArrayToList([]int{1, 1, 2, 2, 3}),
		},
		{
			"test 7",
			args{
				ConvertArrayToList([]int{1, 2, 3}),
				ConvertArrayToList([]int{3, 4}),
			},
			ConvertArrayToList([]int{1, 2, 3, 3, 4}),
		},
		{
			"test 8",
			args{
				ConvertArrayToList([]int{1, 2, 3}),
				ConvertArrayToList([]int{4}),
			},
			ConvertArrayToList([]int{1, 2, 3, 4}),
		},

		{
			"test 9",
			args{
				ConvertArrayToList([]int{1, 3}),
				ConvertArrayToList([]int{2, 4}),
			},
			ConvertArrayToList([]int{1, 2, 3, 4}),
		},
		{
			"test 10",
			args{
				ConvertArrayToList([]int{2, 4}),
				ConvertArrayToList([]int{1, 3}),
			},
			ConvertArrayToList([]int{1, 2, 3, 4}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
