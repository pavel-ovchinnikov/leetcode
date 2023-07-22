package main

import "fmt"

func findCenter(edges [][]int) int {
	if (edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1]) &&
		(edges[0][0] == edges[2][0] || edges[0][0] == edges[2][1]) {
		return edges[0][0]
	}
	return edges[0][1]
}

func main() {
	fmt.Println(findCenter([][]int{
		{1, 2},
		{2, 3},
		{4, 2}}))
}
