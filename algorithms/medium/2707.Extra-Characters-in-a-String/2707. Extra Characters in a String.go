package main

import "fmt"

func minExtraChar(s string, dictionary []string) int {
	dct := make(map[string]bool)
	for _, word := range dictionary {
		dct[word] = true
	}

	n := len(s) + 1
	dp := make([]int, n)
	for i := range dp {
		dp[i] = n
	}
	dp[0] = 0

	for i := 1; i <= len(s); i++ {
		dp[i] = dp[i-1] + 1
		for j := 0; j <= i; j++ {
			if dct[s[i-j:i]] {
				dp[i] = min(dp[i], dp[i-j])
			}
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minExtraChar("sayhelloworld", []string{"hello", "world"}))
}
