package main

import (
	"fmt"
	"strconv"
)

func countSeniors(details []string) int {
	res := 0
	n := 15
	for _, val := range details {
		age, _ := strconv.Atoi(val[n-4 : n-2])
		if age > 60 {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(countSeniors([]string{"7868190130M7522", "5303914400F9211", "9273338290F4010"}))
}
