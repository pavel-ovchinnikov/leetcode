package main

import "fmt"

func countPairs(nums []int, k int) int {
	m := make(map[int][]int, 0)
	res := 0
	for i, v := range nums {
		if idx, ok := m[v]; ok {
			for _, v1 := range idx {
				if v1*i%k == 0 {
					res++
				}
			}
		}
		m[v] = append(m[v], i)
	}
	return res
}

// func countPairs(nums []int, k int) int {
//     res := 0

//     for i := 0; i < len(nums)-1; i++ {
//         for j := i + 1; j < len(nums); j++ {
//             if nums[i] == nums[j] && (i*j) % k == 0 {
//                 res++
//             }
//         }
//     }

//     return res
// }

func main() {
	fmt.Println(countPairs([]int{3, 1, 2, 2, 2, 1, 3}, 2))
	fmt.Println(countPairs([]int{1, 2, 3, 4}, 1))
}
