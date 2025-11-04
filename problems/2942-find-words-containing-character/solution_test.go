package findwordscontainingcharacter

import (
	"reflect"
	"testing"
)

func Test_findWordsContaining(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		words []string
		x     byte
		want  []int
	}{
		{
			name:  "Example 1",
			words: []string{"hello", "world", "leetcode"},
			x:     'o',
			want:  []int{0, 1, 2},
		},
		{
			name:  "Example 2",
			words: []string{"aa", "ab", "ac", "ad"},
			x:     'a',
			want:  []int{0, 1, 2, 3},
		}, {
			name:  "Example 3",
			words: []string{"a", "b", "c"},
			x:     'z',
			want:  []int{},
		}, {
			name:  "Example 4",
			words: []string{"leet", "code"},
			x:     'e',
			want:  []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findWordsContaining(tt.words, tt.x)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWordsContaining() = %v, want %v", got, tt.want)
			}
		})
	}
}
