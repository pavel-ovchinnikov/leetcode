package main

import (
	"fmt"
	"sort"
)

func findDifference(nums1 []int, nums2 []int) [][]int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	i1, i2 := 0, 0
	n1, n2 := len(nums1), len(nums2)
	res1, res2 := make([]int, 0, n1), make([]int, 0, n2)

	for i1 < n1 && i2 < n2 {
		if nums1[i1] < nums2[i2] {
			res1 = append(res1, nums1[i1])
			i1 = nextIndex(i1, nums1)
		} else if nums1[i1] > nums2[i2] {
			res2 = append(res2, nums2[i2])
			i2 = nextIndex(i2, nums2)
		} else {
			i1 = nextIndex(i1, nums1)
			i2 = nextIndex(i2, nums2)
		}
	}

	for i1 < n1 {
		res1 = append(res1, nums1[i1])
		i1 = nextIndex(i1, nums1)
	}

	for i2 < n2 {
		res2 = append(res2, nums2[i2])
		i2 = nextIndex(i2, nums2)
	}

	return [][]int{res1, res2}
}

func nextIndex(i int, arr []int) int {
	if i > len(arr) {
		return i
	}
	j := i + 1
	for ; j < len(arr); j++ {
		if arr[i] != arr[j] {
			return j
		}
	}
	return j
}

func main() {
	// fmt.Println(findDifference([]int{1, 2, 3}, []int{2, 4, 6}))
	fmt.Println(findDifference([]int{1, 2, 3, 3, 4, 4, 4, 2, 5, 5}, []int{1, 1, 2, 2}))

	// i := 0
	// nums := []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 5}
	// for i < len(nums) {

	// 	fmt.Println(nums[i])
	// 	i = nextIndex(i, nums)
	// }
}
