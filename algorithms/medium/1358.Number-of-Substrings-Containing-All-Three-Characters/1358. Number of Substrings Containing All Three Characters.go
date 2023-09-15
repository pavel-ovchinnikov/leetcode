package main

import (
	"fmt"
	"math"
)

func numberOfSubstrings(s string) int {
	n := len(s)
	arrA, arrB, arrC := []int{}, []int{}, []int{}
	for i, c := range s {
		switch c {
		case 'a':
			arrA = append(arrA, i)
		case 'b':
			arrB = append(arrB, i)
		case 'c':
			arrC = append(arrC, i)
		}
	}

	res := 0
	indexA, IndexB, indexC := 0, 0, 0
	for indexA < len(arrA) && IndexB < len(arrB) && indexC < len(arrC) {
		maxIndex := max(arrA[indexA], arrB[IndexB], arrC[indexC])
		minIndex := min(arrA[indexA], arrB[IndexB], arrC[indexC])
		res += n - maxIndex
		switch s[minIndex] {
		case 'a':
			indexA++
		case 'b':
			IndexB++
		case 'c':
			indexC++
		}
	}

	return res
}

func max(nums ...int) int {
	maxVal := math.MinInt
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}
	return maxVal
}

func min(nums ...int) int {
	minVal := math.MaxInt
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}

func main() {
	fmt.Println(numberOfSubstrings("abcabc"))
}
