package main

import "fmt"

func restoreString(s string, indices []int) string {
	res := make([]byte, len(s))
	for i, val := range indices {
		res[val] = s[i]
	}
	return string(res)
}

func main() {
	fmt.Println(restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
	fmt.Println(restoreString("abc", []int{0, 1, 2}))
}
