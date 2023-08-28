package main

import "fmt"

func processQueries(queries []int, m int) []int {
	P := make([]int, m)
	res := make([]int, len(queries))

	for i := 0; i < m; i++ {
		P[i] = i + 1
	}

	for i, num := range queries {
		j := 0
		for P[j] != num {
			j++
		}
		res[i] = j
		for ; j > 0; j-- {
			P[j] = P[j-1]
		}
		P[0] = num
	}
	return res
}

func main() {
	fmt.Println(processQueries([]int{3, 1, 2, 1}, 5))
}
