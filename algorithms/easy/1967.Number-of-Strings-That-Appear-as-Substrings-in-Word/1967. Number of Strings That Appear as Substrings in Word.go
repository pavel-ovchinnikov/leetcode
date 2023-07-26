package main

import (
	"fmt"
	"strings"
)

func numOfStrings(patterns []string, word string) int {
	res := 0
	for _, s := range patterns {
		if strings.Contains(word, s) {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(numOfStrings([]string{"a", "abc", "bc", "d"}, "abc"))
	fmt.Println(numOfStrings([]string{"a", "b", "c"}, "aaaaabbbbb"))
	fmt.Println(numOfStrings([]string{"a", "a", "a"}, "ab"))
	fmt.Println(numOfStrings([]string{
		"dojsf", "v", "hetnovaoigv", "ksvqfdojsf",
		"hetnovaoig", "yskjs", "sfr", "msurztkvppptsluk",
		"ndxffbkkvejuakduhdcfsdbgbt", "znhdgtwzbnh", "h",
		"ocaualfiscmbpnfalypmtdppymw", "w", "o", "sfjksvqfdo",
		"odqvzuc", "bozawheajcmlewptgssueylcxhx", "bno",
		"jhmarzsphxduvdktzqjiz", "j", "sfrjhetnov",
		"vxv", "ksvqfd", "ognwvqtadalmav",
		"yqbspvfwmvhycmghabigl", "qyfaiazgdqaw",
		"ojsfrj", "ojsfrjhetn", "fdojs", "do",
		"ovaoig", "k", "vrh", "jsfrjhetnovaoigvgk"},
		"csfjksvqfdojsfrjhetnovaoigvgk"))

}
