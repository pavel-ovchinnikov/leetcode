package twosum

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Test case 1: Basic input", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{"Test case 2: Negative numbers", args{[]int{-3, 4, 3, 90}, 0}, []int{0, 2}},
		{"Test case 3: No solution", args{[]int{1, 2, 3}, 7}, nil},
		{"Test case 4: Empty array", args{[]int{}, 5}, nil},
		{"Test case 5: Single element", args{[]int{5}, 5}, nil},
		{"Test case 6: Large numbers", args{[]int{1000000, 500000, -1500000}, -1000000}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
