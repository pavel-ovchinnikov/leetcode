package main

import (
	"fmt"
	"strconv"
)

func calPoints(operations []string) int {
	stack := []int{}
	for _, oper := range operations {
		switch oper {
		case "C":
			stack = stack[:len(stack)-1]
		case "D":
			a := stack[len(stack)-1]
			stack = append(stack, a*2)
		case "+":
			a := stack[len(stack)-1] + stack[len(stack)-2]
			stack = append(stack, a)
		default:
			a, _ := strconv.Atoi(oper)
			stack = append(stack, a)
		}
	}
	sum := 0
	for _, val := range stack {
		sum += val
	}
	return sum
}

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
}
