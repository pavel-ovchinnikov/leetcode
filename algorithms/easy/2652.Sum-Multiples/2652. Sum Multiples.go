package main

import "fmt"

func sumOfMultiples(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			res += i
		}
	}
	return res
}

// func sumOfMultiples(n int) int {
// 	nDiv3 := (n / 3)
// 	nDiv5 := (n / 5)
// 	nDiv7 := (n / 7)
// 	nDiv105 := (n / 105)
// 	nDiv15 := (n / 15)
// 	nDiv35 := (n / 35)
// 	nDiv21 := (n / 21)

// 	sum3 := 3 * ((nDiv3 * (nDiv3 + 1)) / 2)
// 	sum5 := 5 * ((nDiv5 * (nDiv5 + 1)) / 2)
// 	sum7 := 7 * ((nDiv7 * (nDiv7 + 1)) / 2)
// 	sum105 := 105 * ((nDiv105 * (nDiv105 + 1)) / 2)
// 	sum15 := 15 * ((nDiv15 * (nDiv15 + 1)) / 2)
// 	sum35 := 35 * ((nDiv35 * (nDiv35 + 1)) / 2)
// 	sum21 := 21 * ((nDiv21 * (nDiv21 + 1)) / 2)

// 	return sum3 + sum5 + sum7 - sum21 - sum15 - sum35 + sum105
// }

func main() {
	fmt.Println(sumOfMultiples(7))
	fmt.Println(sumOfMultiples(10))
	fmt.Println(sumOfMultiples(9))
}
