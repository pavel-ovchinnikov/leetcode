package main

import (
	"fmt"
)

// func canBeEqual(target []int, arr []int) bool {
// 	sort.Ints(target)
// 	sort.Ints(arr)
// 	for i := 0; i < len(target); i++ {
// 		if target[i] != arr[i] {
// 			return false
// 		}
// 	}
// 	return true
// }

func canBeEqual(target []int, arr []int) bool {
	hash := make(map[int]int)

	for i := 0; i < len(target); i++ {
		hash[target[i]]++
		hash[arr[i]]--
		if hash[target[i]] == 0 {
			delete(hash, target[i])
		}
		if hash[arr[i]] == 0 {
			delete(hash, arr[i])
		}
	}

	return len(hash) == 0
}

// func canBeEqual(target []int, arr []int) bool {
//     var values [2][1001]int
//     for _, v := range target {
//         values[0][v]++
//     }
//     for _, v := range arr {
//         values[1][v]++
//     }
//     return values[0] == values[1]
// }

func main() {
	fmt.Println(canBeEqual([]int{1, 2, 3, 4}, []int{2, 4, 1, 3}))
}
