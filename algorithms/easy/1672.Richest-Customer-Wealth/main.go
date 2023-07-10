package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, a := range accounts {
		current := 0
		for _, v := range a {
			current += v
		}
		if max < current {
			max = current
		}
	}
	return max
}

func main() {
	fmt.Println(maximumWealth([][]int{
		{1, 2, 3},
		{3, 2, 1},
	}))
}
