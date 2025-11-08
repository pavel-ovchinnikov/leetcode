package rearrangearrayelementsbysign

import (
	"reflect"
	"testing"
)

func Test_rearrangeArray(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{3, 1, -2, -5, 2, -4},
			want: []int{3, -2, 1, -5, 2, -4},
		},
		{
			name: "Example 2",
			nums: []int{-1, 1},
			want: []int{1, -1},
		}, {
			name: "Custom 4",
			nums: []int{1, -1, 2, -2, 3, -3, 4, -4},
			want: []int{1, -1, 2, -2, 3, -3, 4, -4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rearrangeArray(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rearrangeArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
