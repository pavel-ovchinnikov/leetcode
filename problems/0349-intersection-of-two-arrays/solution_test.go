package task0349

import (
	"reflect"
	"testing"
)

func Test_intersection(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test 01", args{[]int{1, 2, 2, 1}, []int{2, 2}}, []int{2}},
		{"test 02", args{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}}, []int{4, 9}},
		{"test 03", args{[]int{1, 2, 3}, []int{4, 5, 6}}, []int{}},
		{"test 04", args{[]int{}, []int{1}}, []int{}},
		{"test 05", args{[]int{1}, []int{1}}, []int{1}},
		{"test 06", args{[]int{1, 1, 1, 1}, []int{1, 1}}, []int{1}},
		{"test 07", args{[]int{1, 3, 5}, []int{2, 4, 3, 1}}, []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
