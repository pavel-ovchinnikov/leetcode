package main

import "fmt"

func rowAndMaximumOnes(mat [][]int) []int {
	res := []int{0, -1}
	for i := 0; i < len(mat); i++ {
		sum := sum(mat[i])
		if sum > res[1] {
			res[0] = i
			res[1] = sum
		}
	}
	return res
}

func sum(arr []int) int {
	sum := 0
	for j := 0; j < len(arr); j++ {
		sum += arr[j]
	}
	return sum
}

func main() {
	fmt.Println(rowAndMaximumOnes([][]int{
		{0, 0},
		{1, 1},
		{0, 0}}))
}
