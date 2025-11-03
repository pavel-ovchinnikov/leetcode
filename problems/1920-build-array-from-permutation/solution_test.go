package buildarrayfrompermutation

import (
	"reflect"
	"testing"
)

func Test_buildArray(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{0, 2, 1, 5, 3, 4},
			want: []int{0, 1, 2, 4, 5, 3},
		},
		{
			name: "Example 2",
			nums: []int{5, 0, 1, 2, 3, 4},
			want: []int{4, 5, 0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildArray(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
