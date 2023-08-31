package main

import (
	"fmt"
	"sort"
)

func minTaps(n int, ranges []int) int {
	arr := make([][]int, n+1)
	for i := range arr {
		arr[i] = make([]int, 2)
		arr[i][0] = max(0, i-ranges[i])
		arr[i][1] = min(n, i+ranges[i])
	}

	sort.Slice(arr, func(i, j int) bool {
		if arr[i][0] == arr[j][0] {
			return arr[j][1] < arr[i][1]
		}
		return arr[i][0] < arr[j][0]
	})

	if arr[0][0] > 0 {
		return -1
	}
	end := arr[0][1]
	cnt := 0
	for i := range arr {
		cnt++
		if end >= n {
			return cnt
		}
		nextEnd := end
		for j := i + 1; j < len(arr); j++ {
			if arr[j][0] <= end && arr[j][1] > nextEnd {
				i = j
				nextEnd = arr[j][1]
			}
		}
		if nextEnd == end {
			return -1
		}
		end = nextEnd
		i--
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
	fmt.Println(minTaps(5, []int{3, 0, 1, 1, 0, 0}))

}
