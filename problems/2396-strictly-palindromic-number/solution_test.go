package template

import "testing"

func Test_isStrictlyPalindromic(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want bool
	}{
		{
			name: "Example 1",
			n:    4,
			want: false,
		},
		{
			name: "Example 2",
			n:    5,
			want: false,
		},
		{
			name: "Example 3",
			n:    6,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isStrictlyPalindromic(tt.n)
			if got != tt.want {
				t.Errorf("isStrictlyPalindromic() = %v, want %v", got, tt.want)
			}
		})
	}
}
