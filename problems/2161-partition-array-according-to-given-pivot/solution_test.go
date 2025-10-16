package partitionarrayaccordingtogivenpivot

import (
	"reflect"
	"testing"
)

func Test_pivotArray(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums  []int
		pivot int
		want  []int
	}{
		{
			name:  "Example 1",
			nums:  []int{9, 12, 5, 10, 14, 3, 10},
			pivot: 10,
			want:  []int{9, 5, 3, 10, 10, 12, 14},
		},
		{
			name:  "Example 2",
			nums:  []int{-3, 4, 3, 2},
			pivot: 2,
			want:  []int{-3, 2, 4, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pivotArray(tt.nums, tt.pivot)
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pivotArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
