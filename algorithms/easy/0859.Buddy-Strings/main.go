package main

import "fmt"

func buddyStrings(s string, goal string) bool {
	n := len(s)
	if n != len(goal) {
		return false
	}

	frequency := make(map[byte]int)
	count := 0
	for i := 0; i < n; i++ {
		if s[i] != goal[i] {
			count++
			if count > 2 {
				return false
			}
		}

		frequency[s[i]]++
		frequency[goal[i]]--
	}

	if s == goal && len(frequency) == n {
		return false
	}

	for _, v := range frequency {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(buddyStrings("aa", "aa"))
	fmt.Println(buddyStrings("ab", "ba"))
	fmt.Println(buddyStrings("ab", "ab"))
}
