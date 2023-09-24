package main

import (
	"fmt"
	"strconv"
)

func findComplement(num int) int {
	d := strconv.FormatInt(int64(num), 2)
	s := make([]byte, len(d))
	for i := 0; i < len(d); i++ {
		if d[i] == '1' {
			s[i] = '0'
		} else {
			s[i] = '1'
		}
	}
	m, _ := strconv.ParseInt(string(s), 2, 64)
	return int(m)
}

func main() {
	fmt.Println(findComplement(5))
}
