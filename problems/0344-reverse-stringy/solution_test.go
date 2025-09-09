package reversestringy

import "testing"

func Test_reverseString(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s []byte
	}{
		{
			name: "example 1",
			s:    []byte{'h', 'e', 'l', 'l', 'o'},
		},
		{
			name: "example 2",
			s:    []byte{'H', 'a', 'n', 'n', 'a', 'h'},
		},
		{
			name: "single character",
			s:    []byte{'a'},
		},
		{
			name: "two characters",
			s:    []byte{'a', 'b'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.s)
		})
	}
}
