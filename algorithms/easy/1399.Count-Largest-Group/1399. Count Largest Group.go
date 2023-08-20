package main

import (
	"fmt"
	"math"
)

func countLargestGroup(n int) int {
	m := make(map[int][]int)

	for i := 1; i <= n; i++ {
		sum := sumDigit(i)
		m[sum] = append(m[sum], i)
	}

	max := math.MinInt
	count := 0

	for _, lst := range m {
		if len(lst) > max {
			max = len(lst)
			count = 1
			continue
		}

		if len(lst) == max {
			count++
		}

	}

	return count
}

func sumDigit(num int) int {
	sum := 0

	for num > 0 {
		sum += num % 10
		num /= 10
	}

	return sum
}

func main() {
	fmt.Println(countLargestGroup(13))
}
