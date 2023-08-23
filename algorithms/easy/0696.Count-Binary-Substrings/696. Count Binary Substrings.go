package main

import "fmt"

func countBinarySubstrings(s string) int {
	res, cur, prev := 0, 1, 0

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cur++
		} else {
			res += min(prev, cur)
			prev, cur = cur, 1
		}
	}

	return res + min(prev, cur)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func countBinarySubstrings(s string) int {
// 	res, zero, one := 0, 0, 0
// 	i := 0
// 	for i < len(s) {
// 		one, zero = 0, 0
// 		if s[i] == '0' {
// 			for ; i < len(s) && s[i] == '0'; i++ {
// 				zero++
// 			}
// 			for j := i; j < len(s) && s[j] == '1'; j++ {
// 				one++
// 			}
// 		} else {
// 			for ; i < len(s) && s[i] == '1'; i++ {
// 				one++
// 			}
// 			for j := i; j < len(s) && s[j] == '0'; j++ {
// 				zero++
// 			}
// 		}
// 		res += min(zero, one)
// 	}
// 	return res
// }

func main() {
	fmt.Println(countBinarySubstrings("00110011"))
}
