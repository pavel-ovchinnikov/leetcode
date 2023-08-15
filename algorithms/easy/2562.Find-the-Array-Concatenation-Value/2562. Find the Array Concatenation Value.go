package main

import "fmt"

func findTheArrayConcVal(nums []int) int64 {
	l, r := 0, len(nums)-1
	var res int64
	for l < r {
		res += int64(nums[l])*multCount(nums[r]) + int64(nums[r])
		l++
		r--
	}

	if len(nums)%2 == 1 {
		res += int64(nums[len(nums)/2])
	}

	return res
}

func multCount(num int) int64 {
	mult := int64(1)
	for num != 0 {
		mult *= 10
		num /= 10
	}
	return mult
}

func main() {
	fmt.Println(findTheArrayConcVal([]int{7, 52, 2, 4}))
}
