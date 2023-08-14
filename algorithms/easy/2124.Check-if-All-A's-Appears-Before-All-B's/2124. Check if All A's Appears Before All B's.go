package main

import (
	"fmt"
	"strings"
)

func checkString(s string) bool {
	return strings.Count(s, "ba") == 0
}

func main() {
	fmt.Println(checkString("aaabbb"))
}
