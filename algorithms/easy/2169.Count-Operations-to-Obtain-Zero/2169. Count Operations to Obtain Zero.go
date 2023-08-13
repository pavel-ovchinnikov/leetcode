package main

import "fmt"

func countOperations(num1 int, num2 int) int {
	res := 0
	for num1 != 0 && num2 != 0 {
		fmt.Println(num1, num2)
		if num1 > num2 {
			num1 = num1 - num2
		} else {
			num2 = num2 - num1
		}
		res++
	}
	return res
}

func main() {
	fmt.Println(countOperations(5, 4))
}
