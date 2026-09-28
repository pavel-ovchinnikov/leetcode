package task0049

import (
	"reflect"
	"slices"
	"sort"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{name: "nil input", strs: nil, want: [][]string{}},
		{name: "empty input", strs: []string{}, want: [][]string{}},
		{name: "single empty string", strs: []string{""}, want: [][]string{{""}}},
		{name: "single word", strs: []string{"abc"}, want: [][]string{{"abc"}}},
		{
			name: "multiple groups",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			name: "repeated letters",
			strs: []string{"aab", "aba", "baa", "abb"},
			want: [][]string{{"aab", "aba", "baa"}, {"abb"}},
		},
		{
			name: "duplicate words and empty strings",
			strs: []string{"", "abc", "", "bca", "abc"},
			want: [][]string{{"", ""}, {"abc", "bca", "abc"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.strs)
			if !reflect.DeepEqual(normalizeGroups(got), normalizeGroups(tt.want)) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func normalizeGroups(groups [][]string) [][]string {
	normalized := make([][]string, len(groups))
	for i, group := range groups {
		normalized[i] = append([]string{}, group...)
		sort.Strings(normalized[i])
	}
	sort.Slice(normalized, func(i, j int) bool {
		return slices.Compare(normalized[i], normalized[j]) < 0
	})
	return normalized
}
