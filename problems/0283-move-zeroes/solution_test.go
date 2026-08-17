package task0283

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test 01", args{[]int{0, 1, 0, 3, 12}}, []int{1, 3, 12, 0, 0}},
		{"test 02", args{[]int{0}}, []int{0}},
		{"test 03", args{[]int{1}}, []int{1}},
		{"test 04", args{[]int{1, 0}}, []int{1, 0}},
		{"test 05", args{[]int{0, 0, 1}}, []int{1, 0, 0}},
		{"test 06", args{[]int{0, 0, 0}}, []int{0, 0, 0}},
		{"test 07", args{[]int{1, 2, 3}}, []int{1, 2, 3}},
		{"test 08", args{[]int{0, 0, 0, 1, 2, 3}}, []int{1, 2, 3, 0, 0, 0}},
		{"test 09", args{[]int{1, 0, 2, 0, 3, 0}}, []int{1, 2, 3, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}