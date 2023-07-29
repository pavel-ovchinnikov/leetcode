package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)
	n := len(s)

	for i := 0; i < n; i++ {
		v1, ok1 := m1[s[i]]
		v2, ok2 := m2[t[i]]

		if ok1 && ok2 {
			if v1 != t[i] || v2 != s[i] {
				return false
			}
		} else if !ok1 && !ok2 {
			m1[s[i]] = t[i]
			m2[t[i]] = s[i]
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
}
