package reverseprefixofword

import "testing"

func Test_reversePrefix(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		word string
		ch   byte
		want string
	}{
		{
			name: "example 1",
			word: "abcdefd",
			ch:   'd',
			want: "dcbaefd",
		},
		{
			name: "example 2",
			word: "xyxzxe",
			ch:   'z',
			want: "zxyxxe",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reversePrefix(tt.word, tt.ch)
			if got != tt.want {
				t.Errorf("reversePrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
