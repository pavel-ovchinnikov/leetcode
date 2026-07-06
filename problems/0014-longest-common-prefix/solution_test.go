package tasks0014

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{"test 01", []string{"flower", "flow", "flight"}, "fl"},
		{"test 02", []string{"dog", "racecar", "car"}, ""},
		{"test 03", []string{""}, ""},
		{"test 04", []string{"a"}, "a"},
		{"test 05", []string{"ab", "a"}, "a"},
		{"test 06", []string{"", "b"}, ""},
		{"test 07", []string{"same", "same", "same"}, "same"},
		{"test 08", []string{"interspecies", "interstellar", "interstate"}, "inters"},
		{"test 09", []string{"throne", "throne"}, "throne"},
		{"test 10", []string{"throne", "dungeon"}, ""},
		{"test 11", []string{"a", "a", "a"}, "a"},
		{"test 12", []string{"abc", "abc", "abc"}, "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.strs); got != tt.want {
				t.Errorf("longestCommonPrefix(%v) = %q, want %q", tt.strs, got, tt.want)
			}
		})
	}
}
