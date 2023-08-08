package main

import "fmt"

func countPrefixes(words []string, s string) int {
	res := 0
	for _, prefix := range words {
		if isPrefix(prefix, s) {
			res++
		}
	}
	return res
}

func isPrefix(prefix, word string) bool {
	if len(prefix) > len(word) {
		return false
	}
	for i, c := range prefix {
		if c != rune(word[i]) {
			return false
		}
	}
	return true
}

// func countPrefixes(words []string, s string) int {
//     result :=0

//     for i := range words{
//         if strings.HasPrefix(s, words[i]){
//             result++
//         }
//     }
//     return result
// }

func main() {
	fmt.Println(countPrefixes([]string{"a", "b", "c", "ab", "bc", "abc"}, "abc"))
}
