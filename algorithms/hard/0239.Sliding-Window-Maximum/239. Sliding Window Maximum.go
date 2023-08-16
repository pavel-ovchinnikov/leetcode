package main

import (
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	deque := make([]int, 0)

	for i, num := range nums {
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= num {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}

		if deque[0] == i-k+1 {
			deque = deque[1:]
		}
	}

	return res
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
