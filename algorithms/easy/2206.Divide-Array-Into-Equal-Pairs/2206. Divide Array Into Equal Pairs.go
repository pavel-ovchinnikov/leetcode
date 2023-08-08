package main

import "fmt"

func divideArray(nums []int) bool {
	freq := make([]int, 501)

	for _, v := range nums {
		freq[v]++
	}

	for _, v := range freq {
		if v%2 != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(divideArray([]int{3, 2, 3, 2, 2, 2}))
}
