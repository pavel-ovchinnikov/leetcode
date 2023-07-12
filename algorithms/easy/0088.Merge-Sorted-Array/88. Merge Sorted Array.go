package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i1 := m - 1
	i2 := n - 1

	for k := m + n - 1; i2 >= 0; k-- {
		if i1 >= 0 && nums1[i1] > nums2[i2] {
			nums1[k] = nums1[i1]
			i1--
		} else {
			nums1[k] = nums2[i2]
			i2--
		}
	}
}

func main() {
	m := []int{1, 2, 3, 0, 0, 0}
	n := []int{2, 5, 6}
	merge(m, 3, n, 3)
	fmt.Println(m, n)
}
