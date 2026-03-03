package validparentheses

import "testing"

func Test_isValid(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want bool
	}{
		{name: "test 1", s: "()", want: true},
		{name: "test 2", s: "()[]{}", want: true},
		{name: "test 3", s: "(]", want: false},
		{name: "test 4", s: "([])", want: true},
		{name: "test 5", s: "([)]", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.s)
			if got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
