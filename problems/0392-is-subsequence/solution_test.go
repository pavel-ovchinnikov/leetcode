package tasks0392

import "testing"

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty strings", args{"", ""}, true},
		{"empty s", args{"", "abc"}, true},
		{"empty t", args{"abc", ""}, false},
		{"s longer than t", args{"abc", "ab"}, false},
		{"exact match", args{"abc", "abc"}, true},
		{"subsequence at start", args{"abc", "abcdef"}, true},
		{"subsequence at end", args{"def", "abcdef"}, true},
		{"interleaved subsequence", args{"abc", "ahbgdc"}, true},
		{"not a subsequence", args{"axc", "ahbgdc"}, false},
		{"single char match", args{"a", "a"}, true},
		{"single char no match", args{"a", "b"}, false},
		{"duplicate chars", args{"aa", "aab"}, true},
		{"insufficient chars", args{"aa", "ab"}, false},
		{"repeated char in t", args{"aaa", "aaaaa"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
