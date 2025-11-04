package concatenationofarray

import (
	"reflect"
	"testing"
)

func Test_getConcatenation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 1},
			want: []int{1, 2, 1, 1, 2, 1},
		},
		{
			name: "Example 2",
			nums: []int{1, 3, 2, 1},
			want: []int{1, 3, 2, 1, 1, 3, 2, 1},
		},
		{
			name: "Empty array",
			nums: []int{},
			want: []int{},
		},
		{
			name: "Single element",
			nums: []int{5},
			want: []int{5, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getConcatenation(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getConcatenation() = %v, want %v", got, tt.want)
			}
		})
	}
}
