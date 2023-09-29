package main

import "fmt"

func licenseKeyFormatting(s string, k int) string {
	n := len(s)
	nn := 2 * n
	res := make([]byte, nn)
	idx := nn - 1
	count := 0

	for i := n - 1; i >= 0; i-- {
		if s[i] == '-' {
			continue
		}
		if count == k {
			res[idx] = '-'
			idx--
			count = 0
		}

		if s[i] >= 'a' {
			res[idx] = s[i] - ('a' - 'A')
		} else {
			res[idx] = s[i]
		}

		count++
		idx--
	}

	return string(res[idx+1:])
}

func main() {
	// fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
	// fmt.Println(licenseKeyFormatting("2-5g-3-J", 2))
	fmt.Println(licenseKeyFormatting("aaa", 2))
}
