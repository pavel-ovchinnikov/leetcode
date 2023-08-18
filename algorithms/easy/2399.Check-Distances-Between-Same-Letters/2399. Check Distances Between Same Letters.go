package main

import "fmt"

func checkDistances(s string, distance []int) bool {
	count := make([]int, 26)
	for i := len(s) - 1; i >= 0; i-- {
		index := s[i] - 'a'
		count[index] += 1
		if count[index] == 1 {
			distance[index] -= i
		} else {
			distance[index] += i + 1
		}

		if count[index] == 2 && distance[index] != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkDistances("abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
	fmt.Println(checkDistances("aa", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
}
