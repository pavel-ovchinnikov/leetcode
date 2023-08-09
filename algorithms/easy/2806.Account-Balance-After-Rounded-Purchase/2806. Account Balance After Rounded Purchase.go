package main

import "fmt"

func accountBalanceAfterPurchase(purchaseAmount int) int {
	return 100 - round(purchaseAmount)
}

func round(a int) int {
	diff := a % 10
	if diff == 0 {
		return a
	} else if diff >= 5 {
		return a + 10 - diff
	}
	return a - diff
}

func main() {
	fmt.Println(accountBalanceAfterPurchase(23))
}
