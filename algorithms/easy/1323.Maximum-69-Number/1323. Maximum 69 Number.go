package main

import (
	"fmt"
	"strconv"
	"strings"
)

func maximum69Number(num int) int {
	r, _ := strconv.Atoi(strings.Replace(strconv.Itoa(num), "6", "9", 1))
	return r
}

func main() {
	fmt.Println(maximum69Number(9669))
	fmt.Println(maximum69Number(9996))
	fmt.Println(maximum69Number(9999))
}
