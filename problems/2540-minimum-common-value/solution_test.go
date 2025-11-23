package minimumcommonvalue

import "testing"

func Test_getCommon(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			name:  "Test Case 1",
			nums1: []int{1, 2, 3, 6},
			nums2: []int{2, 3, 4, 5},
			want:  2,
		},
		{
			name:  "Test Case 2",
			nums1: []int{1, 2, 3, 4},
			nums2: []int{5, 6, 7, 8},
			want:  -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getCommon(tt.nums1, tt.nums2)
			if got != tt.want {
				t.Errorf("getCommon() = %v, want %v", got, tt.want)
			}
		})
	}
}
