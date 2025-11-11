package distringmatch

import (
	"reflect"
	"testing"
)

func Test_diStringMatch(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want []int
	}{
		{
			name: "example 1",
			s:    "IDID",
			want: []int{0, 4, 1, 3, 2},
		},
		{
			name: "example 2",
			s:    "III",
			want: []int{0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diStringMatch(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diStringMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
