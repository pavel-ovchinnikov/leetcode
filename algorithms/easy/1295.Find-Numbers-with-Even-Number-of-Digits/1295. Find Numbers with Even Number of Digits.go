package main

import "fmt"

func findNumbers(nums []int) int {
	res := 0
	for _, val := range nums {
		count := 0
		for val > 0 {
			val /= 10
			count++
		}

		if count%2 == 0 {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(findNumbers([]int{555, 901, 482, 1771}))
}
