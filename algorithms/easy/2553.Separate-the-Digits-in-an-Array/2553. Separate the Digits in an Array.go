package main

import "fmt"

func separateDigits(nums []int) []int {
	res := make([]int, 0, len(nums))
	digits := make([]int, 10)
	for _, num := range nums {
		i := 9
		for num != 0 {
			digits[i] = num % 10
			num /= 10
			i--
		}
		res = append(res, digits[i+1:]...)
	}
	return res
}

func main() {
	fmt.Println(separateDigits([]int{13, 25, 83, 77}))
}
