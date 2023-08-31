package main

import "fmt"

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	return NumArray{arr: nums}
}

func (numArray *NumArray) SumRange(left int, right int) int {
	sum := 0
	for _, elem := range numArray.arr[left : right+1] {
		sum += elem
	}
	return sum
}

func main() {
	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(numArray.SumRange(0, 2))
	fmt.Println(numArray.SumRange(2, 5))
	fmt.Println(numArray.SumRange(0, 5))
}
