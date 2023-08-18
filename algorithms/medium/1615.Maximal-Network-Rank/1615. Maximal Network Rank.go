package main

import (
	"fmt"
	"math"
)

func maximalNetworkRank(n int, roads [][]int) int {
	res := math.MinInt
	degree := make([]int, n)
	graph := make([][]bool, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]bool, n+1)
	}

	for _, road := range roads {
		graph[road[0]][road[1]] = true
		graph[road[1]][road[0]] = true

		degree[road[0]] += 1
		degree[road[1]] += 1
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if graph[i][j] {
				res = max(res, degree[i]+degree[j]-1)
			} else {
				res = max(res, degree[i]+degree[j])
			}
		}
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(maximalNetworkRank(
		8,
		[][]int{
			{0, 1},
			{1, 2},
			{2, 3},
			{2, 4},
			{5, 6},
			{5, 7},
		}))
}
