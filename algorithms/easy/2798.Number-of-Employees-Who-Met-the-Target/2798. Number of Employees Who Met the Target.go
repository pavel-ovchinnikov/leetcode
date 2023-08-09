package main

import "fmt"

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	res := 0
	for _, val := range hours {
		if val >= target {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(numberOfEmployeesWhoMetTarget([]int{0, 1, 2, 3, 4}, 3))
}
