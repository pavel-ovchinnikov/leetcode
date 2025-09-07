package flippinganimage

import (
	"reflect"
	"testing"
)

func Test_flipAndInvertImage(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		image [][]int
		want  [][]int
	}{
		{
			name:  "Example 1",
			image: [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
			want:  [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := flipAndInvertImage(tt.image)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipAndInvertImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
