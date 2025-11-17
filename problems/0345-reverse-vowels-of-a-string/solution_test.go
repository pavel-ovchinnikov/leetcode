package reversevowelsofastring

import "testing"

func Test_reverseVowels(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want string
	}{
		{
			name: "Test 1",
			s:    "IceCreAm",
			want: "AceCreIm",
		}, {
			name: "Test 2",
			s:    "leetcode",
			want: "leotcede",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseVowels(tt.s)
			if got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
