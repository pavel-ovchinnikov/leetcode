package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 1; i < len(res)+1; i++ {
		if i%15 == 0 {
			res[i-1] = "FizzBuzz"
		} else if i%5 == 0 {
			res[i-1] = "Buzz"
		} else if i%3 == 0 {
			res[i-1] = "Fizz"
		} else {
			res[i-1] = strconv.Itoa(i)
		}
	}
	return res
}

func main() {
	fmt.Println(fizzBuzz(10))
}
