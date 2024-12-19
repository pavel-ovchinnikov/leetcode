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
