package main

import "fmt"

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	alice, bob := 0, 0
	aliceCandies := map[int]bool{}

	for _, candy := range aliceSizes {
		alice += candy
		aliceCandies[candy] = true
	}

	for _, candy := range bobSizes {
		bob += candy
	}
	delta := (alice - bob) / 2
	for _, b := range bobSizes {
		if aliceCandies[b+delta] {
			return []int{b + delta, b}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(fairCandySwap([]int{1, 1}, []int{2, 2}))
	fmt.Println(fairCandySwap([]int{1, 2}, []int{2, 3}))
	fmt.Println(fairCandySwap([]int{2}, []int{1, 3}))
}
