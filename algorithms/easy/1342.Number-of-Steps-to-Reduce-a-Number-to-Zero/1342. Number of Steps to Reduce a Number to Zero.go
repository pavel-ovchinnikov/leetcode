package main

import "fmt"

func numberOfSteps(num int) int {
	res := 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
		res++
	}
	return res
}

func main() {
	fmt.Println(numberOfSteps(10))
	fmt.Println(numberOfSteps(4))
	fmt.Println(numberOfSteps(8))
	fmt.Println(numberOfSteps(123))
}
