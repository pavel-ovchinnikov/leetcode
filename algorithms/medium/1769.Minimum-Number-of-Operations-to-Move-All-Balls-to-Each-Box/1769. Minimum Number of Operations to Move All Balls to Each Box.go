package main

import "fmt"

func minOperations(boxes string) []int {
	n := len(boxes)
	res := make([]int, n)

	counter1 := 0
	counter2 := 0
	for i := 0; i < n; i++ {
		res[i] += counter2
		if boxes[i] == '1' {
			counter1++
		}
		counter2 += counter1
	}

	counter1 = 0
	counter2 = 0
	for i := n - 1; i >= 0; i-- {
		res[i] += counter2
		if boxes[i] == '1' {
			counter1++
		}
		counter2 += counter1
	}
	return res
}

func main() {
	fmt.Println(minOperations("110"))
}
