package longestcommonprefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test1", args: args{strs: []string{"flower", "flow", "flight"}}, want: "fl"},
		{name: "test2", args: args{strs: []string{"dog", "racecar", "car"}}, want: ""},
		{name: "test3", args: args{strs: []string{"a"}}, want: "a"},
		{name: "test4", args: args{strs: []string{"", "b"}}, want: ""},
		{name: "test5", args: args{strs: []string{"ab", "a"}}, want: "a"},
		{name: "test6", args: args{strs: []string{"c", "c"}}, want: "c"},
		{name: "test7", args: args{strs: []string{"aa", "aa"}}, want: "aa"},
		{name: "test8", args: args{strs: []string{"aa", "a"}}, want: "a"},
		{name: "test9", args: args{strs: []string{"aaa", "aa", "aaa"}}, want: "aa"},
		{name: "test10", args: args{strs: []string{"ab", "a"}}, want: "a"},
		{name: "test11", args: args{strs: []string{"abab", "aba", "abc"}}, want: "ab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
