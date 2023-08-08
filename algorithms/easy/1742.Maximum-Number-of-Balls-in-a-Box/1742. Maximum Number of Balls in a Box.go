package main

import "fmt"

func countBalls(lowLimit int, highLimit int) int {
	arr := make([]int, 45)

	index2box := func(a int) (res int) {
		for a > 0 {
			res += a % 10
			a /= 10
		}
		return
	}
	for i := lowLimit; i <= highLimit; i++ {
		arr[index2box(i)-1]++
	}

	maxVal := 0
	for _, v := range arr {
		if maxVal < v {
			maxVal = v
		}
	}
	return maxVal
}

func main() {
	fmt.Println(countBalls(10, 9999))
}
