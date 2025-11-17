package issubsequence

import "testing"

func Test_isSubsequence(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		t    string
		want bool
	}{
		{
			name: "Example 1",
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		}, {
			name: "Example 2",
			s:    "axc",
			t:    "ahbgdc",
			want: false,
		}, {
			name: "Example 3",
			s:    "c",
			t:    "ahbgdc",
			want: true,
		}, {
			name: "Example 4",
			s:    "ahbgdc",
			t:    "a",
			want: false,
		}, {
			name: "Example 5",
			s:    "axc",
			t:    "ahbgdc",
			want: false,
		}, {
			name: "Example 6",
			s:    "",
			t:    "ahbgdc",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSubsequence(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
