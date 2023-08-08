package main

import "fmt"

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	i1, i2 := 0, 0
	n1, n2 := len(nums1), len(nums2)

	res := [][]int{}

	for i1 < n1 && i2 < n2 {
		switch {
		case nums1[i1][0] == nums2[i2][0]:
			res = append(res, []int{nums1[i1][0], nums1[i1][1] + nums2[i2][1]})
			i1++
			i2++
		case nums1[i1][0] < nums2[i2][0]:
			res = append(res, nums1[i1])
			i1++
		default:
			res = append(res, nums2[i2])
			i2++
		}
	}

	for i1 < n1 {
		res = append(res, nums1[i1])
		i1++
	}

	for i2 < n2 {
		res = append(res, nums2[i2])
		i2++
	}
	return res
}

func main() {
	fmt.Println(mergeArrays(
		[][]int{
			{1, 2},
			{2, 3},
			{4, 5}},
		[][]int{
			{1, 4},
			{3, 2},
			{4, 1},
		}))
}
