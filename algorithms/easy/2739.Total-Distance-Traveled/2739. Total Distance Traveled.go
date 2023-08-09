package main

import "fmt"

func distanceTraveled(mainTank int, additionalTank int) int {
	res := 0
	for mainTank > 0 {
		res++
		mainTank--
		if res != 0 && res%5 == 0 && additionalTank >= 1 {
			mainTank++
			additionalTank--
		}
	}
	return res * 10
}

func main() {
	fmt.Println(distanceTraveled(5, 10))
}
