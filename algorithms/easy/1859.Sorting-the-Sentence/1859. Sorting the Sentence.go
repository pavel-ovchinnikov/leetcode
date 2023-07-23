package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sortSentence(s string) string {
	words := strings.Split(s, " ")
	n := len(words)
	result := make([]string, n)
	var i int
	var x string

	for _, v := range words {
		x = string(v[len(v)-1])
		if strings.Contains(v, x) {
			i, _ = strconv.Atoi(x)
			result[i-1] = v[:len(v)-1]
		}
	}

	return strings.Join(result, " ")
}

func main() {
	fmt.Println(sortSentence("is2 sentence4 This1 a3"))
}
