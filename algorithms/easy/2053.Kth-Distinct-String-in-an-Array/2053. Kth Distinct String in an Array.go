package main

import "fmt"

func kthDistinct(arr []string, k int) string {
	res := ""
	freq := make(map[string]int)

	for _, v := range arr {
		freq[v]++
	}

	for _, v := range arr {
		if freq[v] == 1 {
			k--
		}

		if k == 0 {
			res = v
			break
		}
	}

	return res
}

func main() {
	fmt.Println(kthDistinct([]string{"d", "b", "c", "b", "c", "a"}, 2))
}
