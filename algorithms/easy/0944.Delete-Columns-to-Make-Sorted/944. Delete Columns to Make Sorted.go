package main

import "fmt"

func minDeletionSize(strs []string) int {
	isSorted := func(strs []string, col int) bool {
		for i := 1; i < len(strs); i++ {
			if strs[i][col] < strs[i-1][col] {
				return false
			}
		}
		return true
	}
	res := 0
	for i := 0; i < len(strs[0]); i++ {
		if !isSorted(strs, i) {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(minDeletionSize([]string{"cba", "daf", "ghi"}))
	fmt.Println(minDeletionSize([]string{"a", "b"}))
	fmt.Println(minDeletionSize([]string{"zyx", "wvu", "tsr"}))
}
