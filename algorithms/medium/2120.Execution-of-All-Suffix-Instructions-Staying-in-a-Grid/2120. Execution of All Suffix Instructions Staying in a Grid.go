package main

import "fmt"

func executeInstructions(n int, startPos []int, s string) []int {
	res := []int{}
	count := 0

	for i := 0; i < len(s); i++ {

		count = 0
		x, y := startPos[0], startPos[1]

		for ii := i; ii < len(s); ii++ {
			switch string(s[ii]) {
			case "U":
				x--
			case "L":
				y--
			case "R":
				y++
			case "D":
				x++
			}

			if x < 0 || x >= n || y < 0 || y >= n {
				break
			}

			count++
		}

		res = append(res, count)
	}

	return res
}

func main() {
	fmt.Println(executeInstructions(3, []int{0, 1}, "RRDDLU"))
}
