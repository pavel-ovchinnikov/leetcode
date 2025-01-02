package linkedlist

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test 1", args{[]int{1, 2, 3, 4, 0, 1}}, []int{1, 2, 3, 4, 0, 1}},
		{"test 2", args{[]int{}}, []int{}},
		{"test 3", args{[]int{-1}}, []int{-1}},
		{"test 4", args{[]int{5, 4, 3, 2, 1}}, []int{5, 4, 3, 2, 1}},
		{"test 5", args{[]int{0, 0, 0, 0}}, []int{0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := ConvertArrayToList(tt.args.nums)

			if got := ConvertListToArray(node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Convert... = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	type args struct {
		a *ListNode
		b *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test 1",
			args{
				ConvertArrayToList([]int{}),
				ConvertArrayToList([]int{}),
			},
			true,
		},
		{
			"test 2",
			args{
				ConvertArrayToList([]int{1}),
				ConvertArrayToList([]int{1}),
			},
			true,
		},
		{
			"test 3",
			args{
				ConvertArrayToList([]int{1, 2, 3, 4, 5}),
				ConvertArrayToList([]int{1, 2, 3, 4, 5}),
			},
			true,
		},
		{
			"test 4",
			args{
				ConvertArrayToList([]int{}),
				ConvertArrayToList([]int{1, 2, 3, 4, 5}),
			},
			false,
		},
		{
			"test 5",
			args{
				ConvertArrayToList([]int{1, 2}),
				ConvertArrayToList([]int{1, 2, 3, 4, 5}),
			},
			false,
		},
		{
			"test 6",
			args{
				ConvertArrayToList([]int{1, 2, 3, 4, 5}),
				ConvertArrayToList([]int{1, 2}),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
