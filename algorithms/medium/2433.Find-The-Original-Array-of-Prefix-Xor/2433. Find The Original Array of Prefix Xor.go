package main

import "fmt"

func findArray(pref []int) []int {
	n := len(pref)
	for i := n - 1; i > 0; i-- {
		pref[i] = pref[i] ^ pref[i-1]
	}
	return pref
}

func main() {
	fmt.Println(findArray([]int{5, 2, 0, 3, 1}))
}
