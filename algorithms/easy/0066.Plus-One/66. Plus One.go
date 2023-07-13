package main

import "fmt"

func plusOne(digits []int) []int {
	tmp := 1

	for i := len(digits) - 1; i >= 0; i -= 1 {
		tmp += digits[i]

		digits[i] = tmp % 10
		tmp /= 10
		if tmp == 0 {
			break
		}
	}

	if tmp != 0 {
		digits = append([]int{tmp}, digits...)
	}

	return digits
}

func main() {
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{4, 3, 9, 9}))
}
