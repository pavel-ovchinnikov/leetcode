package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	setJewels := make(map[rune]int)
	for _, jewel := range jewels {
		setJewels[jewel] = 0
	}

	result := 0
	for _, stone := range stones {
		if _, ok := setJewels[stone]; ok {
			result += 1
		}
	}
	return result
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
}
