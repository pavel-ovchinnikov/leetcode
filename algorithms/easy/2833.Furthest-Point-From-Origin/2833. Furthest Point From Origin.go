package main

import "fmt"

func furthestDistanceFromOrigin(moves string) int {
	count := 0 // count '_'
	countR := 0
	countL := 0

	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'R':
			countR++
		case 'L':
			countL++
		case '_':
			count++
		}
	}
	return abs(countR-countL) + count
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(furthestDistanceFromOrigin("L_RL__R"))
	fmt.Println(furthestDistanceFromOrigin("_R__LL_"))
	fmt.Println(furthestDistanceFromOrigin("_______"))
}
