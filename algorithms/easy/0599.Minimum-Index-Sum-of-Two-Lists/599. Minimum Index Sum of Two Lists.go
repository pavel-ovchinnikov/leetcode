package main

import (
	"fmt"
	"math"
)

func findRestaurant(list1 []string, list2 []string) []string {
	words := make(map[string]int)
	for i, word := range list1 {
		words[word] = i
	}
	indexSum := math.MaxInt
	res := []string{}
	for i, word := range list2 {
		if idx, ok := words[word]; ok {
			if i+idx == indexSum {
				res = append(res, word)
				continue
			}

			if i+idx < indexSum {
				indexSum = i + idx
				res = []string{word}
				continue
			}

		}
	}
	return res
}

func main() {
	l1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	l2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	fmt.Println(findRestaurant(l1, l2))
}
