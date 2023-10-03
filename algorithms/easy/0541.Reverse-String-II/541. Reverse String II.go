package main

import "fmt"

func reverseStr(s string, k int) string {
	res := []rune(s)
	l, r := 0, k-1
	groupIndex := 0
	for l < r {
		if (r-k+1)%(2*k) == 0 {
			r = Min(r, len(s)-1)
			for l < r {
				res[l], res[r] = res[r], res[l]
				l++
				r--
			}
			groupIndex++
			l = 2 * k * groupIndex
			r = l + k - 1
			if l >= len(s) {
				break
			}
		}
	}

	return string(res)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
}
