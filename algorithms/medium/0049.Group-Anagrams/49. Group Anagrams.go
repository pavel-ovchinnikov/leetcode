package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}

	for _, word := range strs {
		arr := []byte(word)
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
		w := string(arr)
		mp[w] = append(mp[w], word)
	}

	res := make([][]string, 0, len(mp))
	for _, words := range mp {
		res = append(res, words)
	}
	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
