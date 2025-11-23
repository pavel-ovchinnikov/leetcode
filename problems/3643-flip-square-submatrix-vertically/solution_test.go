package flipsquaresubmatrixvertically

import (
	"reflect"
	"testing"
)

func Test_reverseSubmatrix(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		grid [][]int
		x    int
		y    int
		k    int
		want [][]int
	}{
		{
			name: "Example 1",
			grid: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			x: 1,
			y: 1,
			k: 2,
			want: [][]int{
				{1, 2, 3, 4},
				{5, 10, 11, 8},
				{9, 6, 7, 12},
				{13, 14, 15, 16},
			},
		},
		{
			name: "Example 2",
			grid: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			x: 0,
			y: 0,
			k: 3,
			want: [][]int{
				{7, 8, 9},
				{4, 5, 6},
				{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseSubmatrix(tt.grid, tt.x, tt.y, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseSubmatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
