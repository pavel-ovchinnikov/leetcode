package partitionlabels

import (
	"reflect"
	"testing"
)

func Test_partitionLabels(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want []int
	}{
		{
			name: "Example 1",
			s:    "ababcbacadefegdehijhklij",
			want: []int{9, 7, 8},
		},
		{
			name: "Example 2",
			s:    "eccbbbbdec",
			want: []int{10},
		},
		{
			name: "Custom 3",
			s:    "caedbdedda",
			want: []int{1, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partitionLabels(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partitionLabels() = %v, want %v", got, tt.want)
			}
		})
	}
}
