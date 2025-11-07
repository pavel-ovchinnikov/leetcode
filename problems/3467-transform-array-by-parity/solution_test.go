package transformarraybyparity

import (
	"reflect"
	"testing"
)

func Test_transformArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{4, 3, 2, 1},
			want: []int{0, 0, 1, 1},
		},
		{
			name: "Example 2",
			nums: []int{1, 5, 1, 4, 2},
			want: []int{0, 0, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := transformArray(tt.nums)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transformArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
