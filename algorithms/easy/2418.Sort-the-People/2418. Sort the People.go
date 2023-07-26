package main

import (
	"fmt"
	"sort"
)

func sortPeople(names []string, heights []int) []string {
	links := make(map[int]string)
	result := make([]string, 0, len(heights))

	for i := 0; i < len(heights); i++ {
		links[heights[i]] = names[i]
	}

	sort.Slice(heights, func(i, j int) bool {
		return heights[i] > heights[j]
	})

	for i := 0; i < len(heights); i++ {
		result = append(result, links[heights[i]])
	}

	return result
}

func main() {
	fmt.Println(sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170}))
}
