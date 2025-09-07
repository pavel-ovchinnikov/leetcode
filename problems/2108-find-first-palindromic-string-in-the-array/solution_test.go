package findfirstpalindromicstringinthearray

import "testing"

func Test_firstPalindrome(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		words []string
		want  string
	}{
		{
			name:  "Test case 1",
			words: []string{"abc", "car", "ada", "racecar", "cool"},
			want:  "ada",
		},
		{
			name:  "Test case 2",
			words: []string{"notapalindrome", "racecar"},
			want:  "racecar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := firstPalindrome(tt.words)
			if got != tt.want {
				t.Errorf("firstPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
