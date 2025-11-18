package checkifstringisaprefixofarray

import "testing"

func Test_isPrefixString(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s     string
		words []string
		want  bool
	}{
		{
			name:  "example 1",
			s:     "iloveleetcode",
			words: []string{"i", "love", "leetcode", "apples"},
			want:  true,
		},
		{
			name:  "example 2",
			s:     "iloveleetcode",
			words: []string{"apples", "i", "love", "leetcode"},
			want:  false,
		},
		{
			name:  "example 3",
			s:     "abc",
			words: []string{"a", "b", "c"},
			want:  true,
		},
		{
			name:  "example 4",
			s:     "abcd",
			words: []string{"ab", "cd", "ef"},
			want:  true,
		},
		{
			name:  "example 5",
			s:     "abcdxyz",
			words: []string{"ab", "cd", "ef"},
			want:  false,
		},
		{
			name:  "example 6",
			s:     "ccccccccc",
			words: []string{"c", "cc"},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPrefixString(tt.s, tt.words)
			if got != tt.want {
				t.Errorf("isPrefixString() = %v, want %v", got, tt.want)
			}
		})
	}
}
