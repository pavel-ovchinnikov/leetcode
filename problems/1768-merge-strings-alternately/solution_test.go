package mergestringsalternately

import "testing"

func Test_mergeAlternately(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		word1 string
		word2 string
		want  string
	}{
		{
			name:  "Example 1",
			word1: "abc",
			word2: "pqr",
			want:  "apbqcr",
		},
		{
			name:  "Example 2",
			word1: "ab",
			word2: "pqrs",
			want:  "apbqrs",
		},
		{
			name:  "Example 3",
			word1: "abcd",
			word2: "pq",
			want:  "apbqcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeAlternately(tt.word1, tt.word2)
			if got != tt.want {
				t.Errorf("mergeAlternately() = %v, want %v", got, tt.want)
			}
		})
	}
}
