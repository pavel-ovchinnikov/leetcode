package mergetwo2darraysbysummingvalues

import (
	"reflect"
	"testing"
)

func Test_mergeArrays(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums1 [][]int
		nums2 [][]int
		want  [][]int
	}{
		{
			name:  "Example 1",
			nums1: [][]int{{1, 2}, {2, 3}, {4, 5}},
			nums2: [][]int{{1, 4}, {3, 2}, {4, 1}},
			want:  [][]int{{1, 6}, {2, 3}, {3, 2}, {4, 6}},
		},
		{
			name:  "Example 2",
			nums1: [][]int{{2, 4}, {3, 6}, {5, 5}},
			nums2: [][]int{{1, 3}, {4, 3}},
			want:  [][]int{{1, 3}, {2, 4}, {3, 6}, {4, 3}, {5, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeArrays(tt.nums1, tt.nums2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
