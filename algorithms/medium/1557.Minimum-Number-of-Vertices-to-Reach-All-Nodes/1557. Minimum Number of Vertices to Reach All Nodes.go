package main

import "fmt"

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	visited := make([]bool, n)

	for _, edge := range edges {
		visited[edge[1]] = true
	}

	res := make([]int, 0, n-len(visited))
	for i := 0; i < n; i++ {
		if !visited[i] {
			res = append(res, i)
		}
	}

	return res
}

func main() {
	fmt.Println(findSmallestSetOfVertices(6,
		[][]int{
			{0, 1},
			{0, 2},
			{2, 5},
			{3, 4},
			{4, 2}}))
}
