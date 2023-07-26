package main

import (
	"fmt"
)

func countGoodTriplets(arr []int, a int, b int, c int) int {
	n := len(arr)
	res := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			if abs(arr[i]-arr[j]) <= a {
				for k := j + 1; k < n; k++ {
					if abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
						res++
					}
				}
			}
		}
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	fmt.Println(countGoodTriplets([]int{3, 0, 1, 1, 9, 7}, 7, 2, 3))
}
