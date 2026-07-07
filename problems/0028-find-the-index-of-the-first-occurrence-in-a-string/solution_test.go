package tasks0028

import "testing"

func Test_strStr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		want     int
	}{
		{"test 01", "sadbutsad", "sad", 0},
		{"test 02", "leetcode", "leeto", -1},
		{"test 03", "hello", "ll", 2},
		{"test 04", "aaaaa", "bba", -1},
		{"test 05", "a", "a", 0},
		{"test 06", "a", "b", -1},
		{"test 07", "abc", "c", 2},
		{"test 08", "abc", "a", 0},
		{"test 09", "mississippi", "issip", 4},
		{"test 10", "", "", 0},
		{"test 11", "aaa", "a", 0},
		{"test 12", "abcabc", "abc", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.haystack, tt.needle); got != tt.want {
				t.Errorf("strStr(%q, %q) = %v, want %v", tt.haystack, tt.needle, got, tt.want)
			}
		})
	}
}
