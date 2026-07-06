package tasks0020

import "testing"

func Test_isValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"test 01", "()", true},
		{"test 02", "()[]{}", true},
		{"test 03", "(]", false},
		{"test 04", "([)]", false},
		{"test 05", "{[]}", true},
		{"test 06", "", true},
		{"test 07", "[", false},
		{"test 08", "((", false},
		{"test 09", "]", false},
		{"test 10", "([{}])", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.s); got != tt.want {
				t.Errorf("isValid(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
