package main

import "fmt"

// func groupThePeople(groupSizes []int) [][]int {
// 	hashmap := make(map[int][]int)

// 	for index, size := range groupSizes {
// 		hashmap[size] = append(hashmap[size], index)
// 	}

// 	res := [][]int{}
// 	for size, people := range hashmap {
// 		rem := len(people)
// 		for rem > 0 {
// 			res = append(res, people[rem-size:rem])
// 			rem -= size
// 		}
// 	}

// 	return res
// }

func groupThePeople(groupSizes []int) [][]int {
	hash := map[int][]int{}

	for i := 0; i < len(groupSizes); i++ {
		hash[groupSizes[i]] = append(hash[groupSizes[i]], i)
	}

	res := make([][]int, 0, len(hash))
	for size, nums := range hash {
		for i := 0; i < len(nums); i += size {
			res = append(res, nums[i:i+size])
		}
	}

	return res
}

func main() {
	fmt.Println(groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
}
