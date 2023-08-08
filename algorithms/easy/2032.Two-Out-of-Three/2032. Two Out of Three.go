package main

import (
	"fmt"
	"sort"
)

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	sort.Ints(nums3)
	hash := map[int]int{}
	for i := 0; i < len(nums1); {
		hash[nums1[i]]++
		i = next(i, nums1)
	}

	for i := 0; i < len(nums2); {
		hash[nums2[i]]++
		i = next(i, nums2)
	}

	for i := 0; i < len(nums3); {
		hash[nums3[i]]++
		i = next(i, nums3)
	}

	res := []int{}
	for key, val := range hash {
		if val >= 2 {
			res = append(res, key)
		}
	}

	return res
}

func next(index int, arr []int) int {
	if index >= len(arr) {
		return index
	}
	index++
	for index < len(arr) && arr[index-1] == arr[index] {
		index++
	}
	return index
}

// func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) (ans []int) {
// 	get := func(nums []int) (s [101]int) {
// 		for _, v := range nums {
// 			s[v] = 1
// 		}
// 		return
// 	}
// 	s1, s2, s3 := get(nums1), get(nums2), get(nums3)
// 	for i := 1; i <= 100; i++ {
// 		if s1[i]+s2[i]+s3[i] > 1 {
// 			ans = append(ans, i)
// 		}
// 	}
// 	return
// }

func main() {
	fmt.Println(twoOutOfThree(
		[]int{1, 1, 3, 2},
		[]int{1, 1, 3, 2},
		[]int{1, 1, 3, 2},
	))

	fmt.Println(twoOutOfThree(
		[]int{1, 1, 3, 2},
		[]int{2, 3},
		[]int{3},
	))

	fmt.Println(twoOutOfThree(
		[]int{3, 1},
		[]int{2, 3},
		[]int{1, 2},
	))

}
