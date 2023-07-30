package main

import "fmt"

func getRow(rowIndex int) []int {
	lastRow := []int{}
	for i := 0; i < rowIndex+1; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1
		for j := 1; j < i; j++ {
			row[j] = lastRow[j-1] + lastRow[j]
		}
		lastRow = row
	}
	return lastRow
}

func main() {
	fmt.Println(getRow(23))
}
