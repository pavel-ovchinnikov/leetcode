package main

import "fmt"

func prefixesDivBy5(nums []int) []bool {
	ans := make([]bool, len(nums))
	var t int
	for i, num := range nums {
		t = t<<1 + num
		if t%5 == 0 {
			ans[i] = true
			t = 0
		} else {
			t %= 5
		}
	}
	return ans
}

func main() {
	fmt.Println(prefixesDivBy5([]int{0, 1, 1}))
}
