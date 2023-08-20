package main

import "fmt"

func totalMoney(n int) int {
	weeks := n / 7
	leftDays := n % 7
	startMoney := weeks + 1

	//sum of AP series : Sn = (n * [2 * a + (n - 1) * d]) / 2
	completeWeekMoneysum := (weeks * (2*28 + (weeks-1)*7)) / 2
	leftDaysMoneySum := (leftDays * (2*startMoney + (leftDays-1)*1)) / 2

	return completeWeekMoneysum + leftDaysMoneySum
}

func main() {
	fmt.Println(totalMoney(20))
}
