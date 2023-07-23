package main

import "fmt"

func numberOfMatches(n int) int {
	res := 0
	for n > 1 {
		if n%2 == 1 {
			res += n/2 + 1
		} else {
			res += n / 2
		}
		n /= 2
	}

	return res
}

func main() {
	fmt.Println(numberOfMatches(7))
	fmt.Println(numberOfMatches(14))
}
