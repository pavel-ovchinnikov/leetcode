package main

import "fmt"

func bulbSwitch(n int) int {
	arr := make([]bool, n)
	for i := range arr {
		arr[i] = true
	}

	for i := 1; i < n; i++ {
		for j := i; j < n; j += i + 1 {
			arr[j] = !arr[j]
		}
	}

	res := 0
	for i := range arr {
		if arr[i] {
			res++
		}
	}

	return res
}

// func bulbSwitch(n int) int {
//     return int(math.Sqrt(float64(n)))
// }

func main() {
	// fmt.Println(bulbSwitch(8))
	fmt.Println(bulbSwitch(100000000))
}
