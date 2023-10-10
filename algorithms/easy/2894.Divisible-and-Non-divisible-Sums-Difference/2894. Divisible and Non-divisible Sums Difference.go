package main

import "fmt"

// func differenceOfSums(n int, m int) int {
// 	sum1, sum2 := 0, 0
// 	for i := 1; i <= n; i++ {
// 		if i%m != 0 {
// 			sum1 += i
// 		} else {
// 			sum2 += i
// 		}
// 	}
// 	return sum1 - sum2
// }

func differenceOfSums(n int, m int) int {
	sum1 := n * (n + 1) / 2
	sum2 := 2 * m * (n / m) * (n/m + 1) / 2
	return sum1 - sum2
}

func main() {
	fmt.Println(differenceOfSums(10, 3))
}
