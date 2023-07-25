package main

import (
	"fmt"
	"sort"
)

func minMovesToSeat(seats []int, students []int) int {
	result := 0
	sort.Ints(seats)
	sort.Ints(students)
	for i := range seats {
		switch {
		case seats[i] > students[i]:
			result += seats[i] - students[i]
		case seats[i] < students[i]:
			result += students[i] - seats[i]
		}
	}
	return result
}

func main() {
	fmt.Println(minMovesToSeat([]int{3, 1, 5}, []int{2, 7, 4}))
}
