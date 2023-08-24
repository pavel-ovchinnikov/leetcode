package main

import "fmt"

func groupThePeople(groupSizes []int) [][]int {
	hashmap := make(map[int][]int)

	for index, size := range groupSizes {
		hashmap[size] = append(hashmap[size], index)
	}

	res := [][]int{}
	for size, people := range hashmap {
		rem := len(people)
		for rem > 0 {
			res = append(res, people[rem-size:rem])
			rem -= size
		}
	}

	return res
}

func main() {
	fmt.Println(groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
}
