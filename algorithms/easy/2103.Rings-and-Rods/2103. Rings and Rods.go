package main

import "fmt"

func countPoints(rings string) int {
	mm := make(map[int]int)
	size := len(rings)
	res := 0
	i := 0
	for i < size {
		rod := int('9') - int(rings[i+1])
		ring := 0
		switch rings[i] {
		case 'R':
			ring = 1
		case 'G':
			ring = 2
		case 'B':
			ring = 4
		}
		mm[rod] |= ring
		i += 2
	}
	for rod := range mm {
		if mm[rod] == 7 {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(countPoints("B0B6G0R6R0R6G9"))
	fmt.Println(countPoints("B0R0G0R9R0B0G0"))
}
