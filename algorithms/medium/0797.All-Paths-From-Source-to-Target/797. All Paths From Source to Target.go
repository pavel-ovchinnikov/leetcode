package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	res := [][]int{}
	n := len(graph)

	var bfs func(int, []int)
	bfs = func(node int, path []int) {
		if node == n-1 {
			res = append(res, path)
			return
		}

		for _, nextNode := range graph[node] {
			newPath := make([]int, len(path)+1)
			_ = copy(newPath, path)
			newPath[len(newPath)-1] = nextNode
			bfs(nextNode, newPath)
		}
	}

	bfs(0, []int{0})

	return res
}

func main() {
	fmt.Println(allPathsSourceTarget([][]int{
		{4, 3, 1},
		{3, 2, 4},
		{3},
		{4},
		{},
	}))
}
