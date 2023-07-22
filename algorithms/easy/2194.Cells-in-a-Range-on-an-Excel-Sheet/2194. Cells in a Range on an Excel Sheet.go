package main

import (
	"fmt"
	"strconv"
)

func cellsInRange(s string) []string {
	columnStart := rune(s[0])
	rowStart, _ := strconv.Atoi(string(s[1]))
	columnEnd := rune(s[3])
	rowEnd, _ := strconv.Atoi(string(s[4]))

	res := make([]string, 0)

	for rn := columnStart; rn <= columnEnd; rn++ {
		for i := rowStart; i <= rowEnd; i++ {
			res = append(res, fmt.Sprintf("%c%d", rn, i))
		}
	}

	return res
}

func main() {
	fmt.Println(cellsInRange("K1:L2"))
	fmt.Println(cellsInRange("A1:F1"))
}
