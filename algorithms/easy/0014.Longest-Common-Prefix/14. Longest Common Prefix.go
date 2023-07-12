package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	pref := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], pref) {
			pref = pref[:len(pref)-1]
		}
		// fmt.Println(strs[i])
	}
	return pref
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
