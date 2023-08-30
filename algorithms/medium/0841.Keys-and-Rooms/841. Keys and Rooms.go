package main

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	keys := make(map[int]struct{})
	visited := make(map[int]struct{})

	var dfs func(int)
	dfs = func(cur int) {
		visited[cur] = struct{}{}
		for _, room := range rooms[cur] {
			keys[room] = struct{}{}
		}

		for key := range keys {
			if _, ok := visited[key]; !ok {
				dfs(key)
			}
		}
	}

	dfs(0)

	return len(visited) == len(rooms)
}

func main() {
	fmt.Println(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
	fmt.Println(canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
}
