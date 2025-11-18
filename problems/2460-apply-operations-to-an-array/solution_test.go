package applyoperationstoanarray

import (
	"reflect"
	"testing"
)

func Test_applyOperations(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want []int
	}{
		{
			name: "Test 1",
			nums: []int{1, 2, 2, 1, 1, 0},
			want: []int{1, 4, 2, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := applyOperations(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
