package main

import "fmt"

func wateringPlants(plants []int, capacity int) int {
	tank := capacity
	count := 0
	for i, plant := range plants {
		if tank >= plant {
			tank -= plant
		} else {
			tank = capacity - plant
			count += 2 * i
		}
		count++
	}
	return count
}

func main() {
	fmt.Println(wateringPlants([]int{7, 7, 7, 7, 7, 7}, 8))
}
