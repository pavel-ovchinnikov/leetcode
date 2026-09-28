package task0049

import "sort"

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string, len(strs))

	for _, s := range strs {
		key := sortRune(s)
		groups[key] = append(groups[key], s)
	}

	res := make([][]string, 0, len(groups))
	for _, g := range groups {
		res = append(res, g)
	}
	return res
}

func sortRune(a string) string {
	ra := []rune(a)
	sort.Slice(ra, func(i, j int) bool { return ra[i] < ra[j] })
	return string(ra)
}
