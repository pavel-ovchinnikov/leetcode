package main

import "fmt"

func predictPartyVictory(senate string) string {
	dir := []int{}
	rad := []int{}
	n := len(senate)

	for i, c := range senate {
		if c == 'R' {
			rad = append(rad, i)
		} else {
			dir = append(dir, i)
		}
	}

	for len(rad) != 0 && len(dir) != 0 {
		n++
		if rad[0] < dir[0] {
			rad = append(rad, n)
		} else {
			dir = append(dir, n)
		}
		rad = rad[1:]
		dir = dir[1:]
	}

	if len(rad) == 0 {
		return "Dire"
	} else {
		return "Radiant"
	}
}

func main() {
	fmt.Println(predictPartyVictory("RD"))
	fmt.Println(predictPartyVictory("RDD"))
}
