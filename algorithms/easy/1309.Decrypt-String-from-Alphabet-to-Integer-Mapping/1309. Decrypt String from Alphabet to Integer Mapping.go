package main

import (
	"fmt"
	"strconv"
)

func freqAlphabets(s string) string {
	res := make([]byte, 0, len(s))
	n := len(s)
	var num int

	for i := 0; i < n; {
		if i < n-2 && s[i+2] == '#' {
			num, _ = strconv.Atoi(string(s[i]) + string(s[i+1]))
			res = append(res, byte('a'+num-1))
			i += 3
		} else {
			num, _ = strconv.Atoi(string(s[i]))
			res = append(res, byte('a'+num-1))
			i++
		}
	}
	return string(res)
}

func main() {
	fmt.Println(freqAlphabets("10#11#12"))
	fmt.Println(freqAlphabets("1326#"))
	fmt.Println(freqAlphabets("12345"))
}
