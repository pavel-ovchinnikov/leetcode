package sortarraybyparityii

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParityII(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{4, 2, 5, 7},
			want: []int{4, 5, 2, 7},
		},
		{
			name: "Example 2",
			nums: []int{2, 3},
			want: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortArrayByParityII(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParityII() = %v, want %v", got, tt.want)
			}
		})
	}
}
