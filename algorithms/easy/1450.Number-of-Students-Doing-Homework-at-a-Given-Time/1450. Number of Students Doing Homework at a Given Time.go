package main

import "fmt"

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	n := len(startTime)
	res := 0
	for i := 0; i < n; i++ {
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(busyStudent(
		[]int{1, 2, 3},
		[]int{3, 2, 7},
		4))
}
