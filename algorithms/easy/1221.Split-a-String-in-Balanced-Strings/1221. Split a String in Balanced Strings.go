package main

import "fmt"

func balancedStringSplit(s string) int {
	result := 0
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			count++
		} else {
			count--
		}

		if count == 0 {
			result++
		}
	}
	return result
}

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL"))
	fmt.Println(balancedStringSplit("RLRRRLLRLL"))
}
