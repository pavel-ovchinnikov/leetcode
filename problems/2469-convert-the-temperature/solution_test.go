package convertthetemperature

import (
	"testing"

	"github.com/pavel-ovchinnikov/leetcode/tools/slices"
)

func Test_convertTemperature(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		celsius float64
		want    []float64
	}{
		{
			name:    "Example 1",
			celsius: 36.50,
			want:    []float64{309.65, 97.7},
		},
		{
			name:    "Example 2",
			celsius: 122.11,
			want:    []float64{395.26, 251.798},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertTemperature(tt.celsius)
			if !slices.CompareFloatSlices(got, tt.want, 0.01) {
				t.Errorf("convertTemperature() = %v, want %v", got, tt.want)
			}
		})
	}
}
