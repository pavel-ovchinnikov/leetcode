package main

import "fmt"

func minTaps(n int, ranges []int) int {
	arr := make([]int, n+1)
	for i, rng := range ranges {
		l := max(0, i-rng)
		r := i + rng
		arr[l] = max(arr[l], r)
	}

	globalMax := arr[0]
	localMax := arr[0]
	count := 1
	for i := 0; i < n; i++ {
		if i > globalMax {
			return -1
		}

		globalMax = max(globalMax, arr[i])
		if i == localMax {
			localMax = globalMax
			count++
		}
	}
	return count
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(minTaps(5, []int{3, 4, 1, 1, 0, 0}))
	fmt.Println(minTaps(3, []int{0, 0, 0, 0}))
	fmt.Println(minTaps(5, []int{1, 1, 1, 1, 1, 1}))
	fmt.Println(minTaps(10, []int{1, 4, 1, 1, 0, 0, 2, 1, 5, 6, 0}))

}
