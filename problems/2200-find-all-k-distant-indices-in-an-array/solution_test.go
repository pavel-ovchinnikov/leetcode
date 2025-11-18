package findallkdistantindicesinanarray

import (
	"reflect"
	"testing"
)

func Test_findKDistantIndices(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		key  int
		k    int
		want []int
	}{
		{
			name: "example 1",
			nums: []int{3, 4, 9, 1, 3, 9, 5},
			key:  9,
			k:    1,
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "example 2",
			nums: []int{2, 2, 2, 2, 2},
			key:  2,
			k:    2,
			want: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findKDistantIndices(tt.nums, tt.key, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findKDistantIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
