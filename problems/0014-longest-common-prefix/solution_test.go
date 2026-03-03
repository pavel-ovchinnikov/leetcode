package longestcommonprefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		strs []string
		want string
	}{
		{name: "test1", strs: []string{"flower", "flow", "flight"}, want: "fl"},
		{name: "test2", strs: []string{"dog", "racecar", "car"}, want: ""},
		{name: "test3", strs: []string{"a"}, want: "a"},
		{name: "test4", strs: []string{"", "b"}, want: ""},
		{name: "test5", strs: []string{"ab", "a"}, want: "a"},
		{name: "test6", strs: []string{"c", "c"}, want: "c"},
		{name: "test7", strs: []string{"aa", "aa"}, want: "aa"},
		{name: "test8", strs: []string{"aa", "a"}, want: "a"},
		{name: "test9", strs: []string{"aaa", "aa", "aaa"}, want: "aa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefix(tt.strs)
			if got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
