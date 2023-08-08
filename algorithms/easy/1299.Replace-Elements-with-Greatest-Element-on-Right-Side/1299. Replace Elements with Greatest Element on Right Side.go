package main

import "fmt"

func replaceElements(arr []int) []int {
	max := -1
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > max {
			max, arr[i] = arr[i], max
		} else {
			arr[i] = max
		}
	}
	return arr
}

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1}))
}
