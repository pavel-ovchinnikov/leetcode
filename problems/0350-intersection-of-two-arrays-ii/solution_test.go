package intersectionoftwoarraysii

import (
	"reflect"
	"testing"
)

func Test_intersect(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			name:  "Example 1",
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
			want:  []int{2, 2},
		},
		{
			name:  "Example 2",
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
			want:  []int{4, 9},
		},
		{
			name:  "No Intersection",
			nums1: []int{1, 2, 3},
			nums2: []int{4, 5, 6},
			want:  []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intersect(tt.nums1, tt.nums2)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
