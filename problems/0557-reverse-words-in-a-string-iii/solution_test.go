package reversewordsinastringiii

import "testing"

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want string
	}{
		{
			name: "Example 1",
			s:    "Let's take LeetCode contest",
			want: "s'teL ekat edoCteeL tsetnoc",
		},
		{
			name: "Example 2",
			s:    "God Ding",
			want: "doG gniD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseWords(tt.s)
			if got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
