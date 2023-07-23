package main

import "fmt"

func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	res := 0
	for i := 0; i < n; i++ {
		res += ((i+1)*(n-i) + 1) / 2 * arr[i]
	}
	return res
}

func main() {
	fmt.Println(sumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
}
