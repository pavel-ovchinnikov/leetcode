package lexicographicallysmallestpalindrome

import "testing"

func Test_makeSmallestPalindrome(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want string
	}{
		{
			name: "Test 1",
			s:    "egcfe",
			want: "efcfe",
		},
		{
			name: "Test 2",
			s:    "abcd",
			want: "abba",
		},
		{
			name: "Test 3",
			s:    "seven",
			want: "neven",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeSmallestPalindrome(tt.s)
			if got != tt.want {
				t.Errorf("makeSmallestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
