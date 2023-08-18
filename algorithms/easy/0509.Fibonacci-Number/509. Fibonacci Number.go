package main

import "fmt"

func fib(n int) int {
	f0 := 0
	f1 := 1
	if n == 0 {
		return f0
	}

	for n > 1 {
		f0, f1 = f1, f0+f1
		n--
	}
	return f1
}

func main() {
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(5))
}
